package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// GenerateRSAKeyPair generates a private and public key pair, each 4096 bits
// long, using the RSA algorithm.
func GenerateRSAKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, _ := rsa.GenerateKey(rand.Reader, 4096)
	return privkey, &privkey.PublicKey
}

// ExportRSAPrivateKeyAsPEMStr will take an RSA-generated private key and convert it
// into a PEM-formated string.
func ExportRSAPrivateKeyAsPEMStr(privkey *rsa.PrivateKey) string {
	privkeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privkey),
		},
	)

	return string(privkeyPEM)
}

// ParseRSAPrivateKeyFromPEMStr will take a pem-formated RSA private key and
// parse it into an *rsa.PrivateKey type.
func ParseRSAPrivateKeyFromPEMStr(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

// ExportRSAPublicKeyAsPEMStr will take an RSA-generated public key and
// convert it into a PEM-formated string.
func ExportRSAPublicKeyAsPEMStr(pubkey *rsa.PublicKey) (string, error) {
	pubkeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLICK KEY",
			Bytes: pubkeyBytes,
		},
	)

	return string(pubkeyPEM), nil
}

// ParseRSAPublickKeyFomPEMStr will take a pem-fromated RSA public key and
// parse it into an *rsa.PublicKey type.
func ParseRSAPublicKeyFromPEMStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break
	}
	return nil, errors.New("Key is not RSA")
}

func main() {
	// Create key pair.
	priv, pub := GenerateRSAKeyPair()

	// Export keys to pem string.
	privPem := ExportRSAPrivateKeyAsPEMStr(priv)
	pubPem, err := ExportRSAPublicKeyAsPEMStr(pub)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Import keys from pem string.
	privParsed, err := ParseRSAPrivateKeyFromPEMStr(privPem)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pubParsed, err := ParseRSAPublicKeyFromPEMStr(pubPem)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Export the newly imported keys
	priv_parsed_pem := ExportRSAPrivateKeyAsPEMStr(privParsed)
	pub_parsed_pem, err := ExportRSAPublicKeyAsPEMStr(pubParsed)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

	fmt.Println(priv_parsed_pem)
	fmt.Println(pub_parsed_pem)

	// Check that the exported/imported keys match the original keys
	if privPem != priv_parsed_pem || pubPem != pub_parsed_pem {
		fmt.Println("Failure: Export and Import did not result in same Keys")
	} else {
		fmt.Println("Success")
	}
}
