package main

import (
	"io/ioutil"

	crypto "github.com/libp2p/go-libp2p-crypto"
	"github.com/mr-tron/base58"
)

func loadIdentity(filepath string) (crypto.PrivKey, error) {

	keyBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	prvKey, err := UnmarshalPrvKey(string(keyBytes))
	if err != nil {
		return nil, err
	}

	return prvKey, nil
}

//UnmarshalPrvKey from base58 string to private key
func UnmarshalPrvKey(prvKey string) (crypto.PrivKey, error) {
	decoded, err := base58.Decode(prvKey)
	if err != nil {
		return nil, err
	}
	unmarshalKey, err := crypto.UnmarshalPrivateKey(decoded)
	if err != nil {
		return nil, err
	}
	return unmarshalKey, nil
}
