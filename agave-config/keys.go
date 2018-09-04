package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// keysResponseError holds the values for a failed request to the key's api
// endpoint.
type keysResponseError struct {
	Fault  keysResponseErrorContent `json:"fault"`
	Status string
}

// keysResponseErrorContent represents the "fault" field from a failed response
// to the keys endpoint.
type keysResponseErrorContent struct {
	Code        int64  `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

// Error formats an error message for a failed response to the keys endpoint.
func (e *keysResponseError) Error() string {
	return fmt.Sprintf("%s: %s(%d) - %s",
		e.Status, e.Fault.Message, e.Fault.Code, e.Fault.Description)
}

// GetUserPubKeys makes a request for plain-text public keys for a given user.
func (c *Configurations) GetUserPubKeys(user string) error {
	// Keys endpoint.
	keysEndpoint := c.BaseUrl + "/keys/v2/" + user + "/text"

	// Make request.
	req, err := http.NewRequest("GET", keysEndpoint, nil)
	if err != nil {
		fmt.Printf("Error building request: %s\n", err)
		return err
	}

	// Set request headers.
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)

	// Create HTTP client with timeout of 10s.
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// Make HTTP request.
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %s\n", err)
		return err
	}
	defer resp.Body.Close()

	// Check if request was successful.
	if resp.StatusCode == http.StatusOK {
		// Pass public keys to stdout.
		reader := bufio.NewReader(resp.Body)
		for {
			key, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}

				fmt.Printf("Error while reading: %s\n", err)
				return err
			}
			fmt.Printf("%s", key)
		}
	} else { // API call returned an error.
		var failedResp keysResponseError
		if err := json.NewDecoder(resp.Body).Decode(&failedResp); err != nil {
			fmt.Printf("Error decoding failed response's body: %s\n", err)
			return err
		}
		failedResp.Status = resp.Status
		return &failedResp
	}

	return nil
}
