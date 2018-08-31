package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// agaveConfigs stores the credentials necessary to interact with the Agave API.
type agaveConfigs struct {
	TenantId     string `json:"tenantid"`
	BaseUrl      string `json:"baseurl"`
	ApiSecret    string `json:"apisecret"`
	ApiKey       string `json:"apikey"`
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    string `json:"created_at"`
	ExpiresIn    string `json:"expires_in"`
	ExpiresAt    string `json:"expires_at"`
}

type refreshToken struct {
	Scope        string `json:"scope"`
	TokeType     string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// parseAgaveConfig reads the agave credentials file and stores it as an
// agaveConfigs struct.
func parseAgaveConfig(filename string) (*agaveConfigs, error) {
	configFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening JSON file: %s\n", err)
		return nil, err
	}
	defer configFile.Close()

	var agaveConf agaveConfigs
	if err := json.NewDecoder(configFile).Decode(&agaveConf); err != nil {
		fmt.Printf("Error decoding JSON file: %s\n", err)
		return nil, err
	}

	return &agaveConf, nil
}

func updateAgaveConfig(filename string, agaveConf *agaveConfigs) error {
	configFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("Error opening file for read and write: %s\n", err)
		return err
	}
	defer configFile.Close()

	encoder := json.NewEncoder(configFile)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(&agaveConf); err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return err
	}

	return err
}

func main() {
	agaveConf, err := parseAgaveConfig("config.json")
	if err != nil {
		fmt.Println(err)
	}

	endpoint := "https://api.tacc.utexas.edu//token"
	v := url.Values{}
	v.Set("grant_type", "refresh_token")
	v.Set("scope", "PRODUCTION")
	v.Set("refresh_token", agaveConf.RefreshToken)
	data := v.Encode()
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data))
	req.SetBasicAuth(agaveConf.ApiKey, agaveConf.ApiSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		fmt.Printf("Error building request: %s\n", err)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
	}
	defer resp.Body.Close()

	fmt.Printf("response Status: %v\n", resp.Status)
	fmt.Printf("response Headers: %s\n", resp.Header["Date"])

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	fmt.Println("Response: ", string(bodyBytes))

	if resp.StatusCode == http.StatusOK {
		var refreshedToken refreshToken
		if err := json.NewDecoder(resp.Body).Decode(&refreshedToken); err != nil {
			fmt.Printf("Error decoding response body: %s\n", err)
			os.Exit(1)
		}

		agaveConf.RefreshToken = refreshedToken.RefreshToken
		agaveConf.AccessToken = refreshedToken.AccessToken

		if err := updateAgaveConfig("config.json", agaveConf); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
