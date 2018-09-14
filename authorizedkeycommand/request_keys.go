package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// GetUserPubKeys makes a request for plain-text public keys for a given user.
func GetUserPubKeys(user string) error {
	// Keys endpoint.
	reqEndpoint := keysEndpoint + user + "/text"

	// Make request.
	req, err := http.NewRequest("GET", reqEndpoint, nil)
	if err != nil {
		fmt.Printf("Error building request: %s\n", err)
		return err
	}

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

				return err
			}
			fmt.Printf("%s", key)
		}
	} else { // API call returned an error.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return nil
}
