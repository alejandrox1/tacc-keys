package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// Request payload
type Payload struct {
	KeyValue string        `json:"key_value"`
	Tags     []PayloadTags `json:"tags"`
}

type PayloadTags struct {
	Purpose string `json:"name"`
}

// PostUserPubKey posts a user's public key to the keys server.
func (c *Configurations) PostUserPubKey(user string, pubkey string) error {
	// Keys endpoint.
	keysEndpoint := c.BaseUrl + "/keys/v2/" + user

	// Request payload.
	data := Payload{
		KeyValue: pubkey,
		Tags:     []PayloadTags{{Purpose: "keyservice-test"}},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Make request.
	req, err := http.NewRequest("POST", keysEndpoint, bytes.NewReader(payloadBytes))
	if err != nil {
		fmt.Printf("Error building request: %s\n", err)
		return err
	}

	// Set request headers.
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

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
		// Print response to stdout.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(body))
	} else { // API call returned an error.
		var failedResp keysResponseError
		if err := json.NewDecoder(resp.Body).Decode(&failedResp); err != nil {
			fmt.Printf("Error decoding failed response's body: %s\n", err)
			fmt.Printf("%+v\n", data)
			return err
		}
		failedResp.Status = resp.Status
		return &failedResp
	}

	return nil
}
