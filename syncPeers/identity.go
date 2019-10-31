package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"

	crypto "github.com/libp2p/go-libp2p-crypto"
)

//DecodeHexToPrvKey from hex encode to private key
func DecodeHexToPrvKey(prvKey []byte) (crypto.PrivKey, error) {
	hexDecodeKey := make([]byte, hex.DecodedLen(len(prvKey)))
	_, err := hex.Decode(hexDecodeKey, prvKey)
	if err != nil {
		return nil, err
	}

	unmarshalKey, err := crypto.UnmarshalPrivateKey(hexDecodeKey)
	if err != nil {
		return nil, err
	}
	return unmarshalKey, nil
}

func loadIdentity(filepath string) (crypto.PrivKey, error) {

	keyBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\x1b[31mKey=%s\x1b[0m\n", string(keyBytes))
	prvKey, err := DecodeHexToPrvKey(keyBytes)
	if err != nil {
		return nil, err
	}

	return prvKey, nil
}
