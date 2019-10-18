package main

import (
	"crypto/rand"
	"flag"
	"io/ioutil"
	"os"
	"path"
	"strconv"

	crypto "github.com/libp2p/go-libp2p-crypto"
	"github.com/mr-tron/base58"
)

var prefixKey string
var countKey int

func main() {
	flag.StringVar(&prefixKey, "prefix", "peer", "node private key path")
	flag.IntVar(&countKey, "count", 4, "number of key to generate")
	flag.Parse()

	for i := 1; i <= countKey; i++ {
		saveGenKey(i, prefixKey)
	}

}

func saveGenKey(count int, prefix string) error {
	prv, err := randKey()
	if err != nil {
		return err
	}

	keyfile := path.Join(os.Getenv("PWD"), "key", prefix+strconv.Itoa(count)+".prv")

	encodedKey, err := marshalPrvKey(prv)
	if err := ioutil.WriteFile(keyfile, []byte(encodedKey), 0644); err != nil {
		return err
	}

	return nil
}

func randKey() (crypto.PrivKey, error) {
	r := rand.Reader
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)

	if err != nil {
		return nil, err
	}
	return prvKey, nil
}
func marshalPrvKey(prvKey crypto.PrivKey) (string, error) {
	marshalKey, err := crypto.MarshalPrivateKey(prvKey)
	if err != nil {
		return "", err
	}
	encoded := base58.Encode(marshalKey)
	return encoded, nil
}
