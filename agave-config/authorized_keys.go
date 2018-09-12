package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ParseAuthroizedKeys(user string) error {
	// Open authorized keys file.
	authorizedKeys, err := os.Open("/home/" + user + "/.ssh/authorized_keys")
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
