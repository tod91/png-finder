package checker

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
)

func Equal(src, dest string) bool {
	return getHash(src) == getHash(dest)
}

var cache = map[string]string{}

func getHash(file string) string {
	if r, ok := cache[file]; ok {
		return r
	}

	hasher := sha256.New()
	s, err := ioutil.ReadFile(file)
	hasher.Write(s)
	if err != nil {
		log.Fatal(err)
	}

	cache[file] = hex.EncodeToString(hasher.Sum(nil))
	return cache[file]
}
