package main

import (
	"fmt"
	"encoding/hex"
	"log"
)

func main() {
	m := []byte("attack at dawn")
	c, err := hex.DecodeString("6c73d5240a948c86981bc294814d")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PT: %b\n", m)
	fmt.Printf("CT: %b\n", c)
	key := xorBytes(m, c)
	fmt.Println(key)
	fmt.Println(hex.EncodeToString(xorBytes(m, key)))
	fmt.Println(string(xorBytes(c, key)))
}

func xorBytes(b1 []byte, bmore ...[]byte) []byte {
	for _, m := range bmore {
		if len(b1) != len(m) {
			panic("length mismatch")
		}
	}

	rv := make([]byte, len(b1))

	for i := range b1 {
		for _, m := range bmore {
			rv[i] = b1[i] ^ m[i]
		}
	}

	return rv
}
