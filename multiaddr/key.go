package main

import (
	"crypto/rand"
	"encoding/hex"

	crypto "github.com/libp2p/go-libp2p-crypto"
)

func GenRandPrvKey() (crypto.PrivKey, error) {
	r := rand.Reader
	prvKey, _, err := crypto.GenerateEd25519Key(r)
	if err != nil {
		return nil, err
	}
	return prvKey, nil
}

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

// EncodePrvKeyToHex  from hex encoded string to private key
func EncodePrvKeyToHex(prvKey crypto.PrivKey) ([]byte, error) {
	marshalKey, err := crypto.MarshalPrivateKey(prvKey)
	if err != nil {
		return nil, err
	}
	hexEncodeKey := make([]byte, hex.EncodedLen(len(marshalKey)))
	hex.Encode(hexEncodeKey, marshalKey)
	return hexEncodeKey, nil
}
