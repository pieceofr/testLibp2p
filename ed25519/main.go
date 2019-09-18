package main

import (
	"crypto/rand"
	"fmt"

	crypto "github.com/libp2p/go-libp2p-crypto"
)

func main() {
	prvKey, err := randKey()
	if err != nil {
		panic(err)
	}
	prvKeyRaw, err := prvKey.Raw()
	if err != nil {
		panic(err)
	}
	fmt.Println("private key raw len=", len(prvKeyRaw))

	prvKeyByte, err := prvKey.Bytes()
	if err != nil {
		panic(err)
	}
	fmt.Println("private key proto buff byte len=", len(prvKeyByte))

	pubKey := prvKey.GetPublic()
	pubKeyRaw, err := pubKey.Raw()
	if err != nil {
		panic(err)
	}
	fmt.Println("public  key raw len=", len(pubKeyRaw))

	pubKeyBytes, err := pubKey.Bytes()
	if err != nil {
		panic(err)
	}
	fmt.Println("public  key bytes len=", len(pubKeyBytes))

}

func randKey() (crypto.PrivKey, error) {
	r := rand.Reader
	prvKey, _, err := crypto.GenerateEd25519Key(r)
	if err != nil {
		return nil, err
	}
	return prvKey, nil
}
