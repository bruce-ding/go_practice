package main

import (
	cryptorand "crypto/rand"
	"encoding/base64"
	"fmt"
	mathrand "math/rand"
)

func RandomString() string {
	bytes := make([]byte, 12)

	_, err := cryptorand.Read(bytes)

	if err != nil {
		mathrand.Read(bytes)
	}

	return base64.RawURLEncoding.EncodeToString(bytes)
}

func main() {
	str := RandomString()
	fmt.Println(len(str))
}
