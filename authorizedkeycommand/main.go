/*
   This program is designed to be used as an sshd AuthorizedKeysCommand.
   sshd requires the command to output public keys to stdout and for each key
   to be delimited by a newline character.
*/
package main

import (
	"fmt"
	"os"
)

// keysEndpoint is the uri for the service to which we will request ssh public
// keys from.
const keysEndpoint string = "http://keys.tacc.cloud/keys/"

func main() {
	// Check that a command line argument was passed. We expect this argument
	// to be the user trying to login.
	if len(os.Args) <= 1 {
		os.Exit(1)
	}

	username := os.Args[1]

	if err := GetUserPubKeys(username); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
