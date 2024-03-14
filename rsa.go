package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func generateRSAKeys(bits int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, bits)
}

func savePrivateKey(key *rsa.PrivateKey, filename string) error {
	privBytes := x509.MarshalPKCS1PrivateKey(key)
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = pem.Encode(file, pemBlock)
	if err != nil {
		return err
	}

	return nil
}

func savePublicKey(pub *rsa.PublicKey, filename string) error {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err
	}

	pemBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = pem.Encode(file, pemBlock)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// generate 1024 bits RSA Keys
	privateKey, err := generateRSAKeys(1024)
	if err != nil {
		fmt.Println("Generate RSA Keys Error:", err)
		return
	}

	// save private key to file (private.pem)
	err = savePrivateKey(privateKey, "private.pem")
	if err != nil {
		fmt.Println("Save Private Key To File Error:", err)
		return
	}

	// get public key
	publicKey := &privateKey.PublicKey

	// save public key to file (public.pem)
	err = savePublicKey(publicKey, "public.pem")
	if err != nil {
		fmt.Println("Save Public Key To File Error:", err)
		return
	}

	fmt.Println("Generate Success!")
}
