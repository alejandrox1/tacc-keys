package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)


// ParseAuthroizedKeys will read ~/.ssh/authorized_keys and output the file
// content's to stdout.
func ParseAuthroizedKeys(user string) error {
	authorizedKeysFile := "/home/" + user + "/.ssh/authorized_keys"
	// Check file ~/.ssh/authorized_keys file exists.
	if _, err := os.Stat(authorizedKeysFile); os.IsNotExist(err) {
		return nil
	}

	// Open authorized keys file.
	authorizedKeys, err := os.Open(authorizedKeysFile)
	if err != nil {
		return err
	}
	defer authorizedKeys.Close()

	reader := bufio.NewReader(authorizedKeys)
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

	return nil
}
