package main

import (
	"fmt"
	"encoding/hex"
	"log"
)

func main() {
	m := []byte("attack at dawn")
	m2 := []byte("attack at dusk")
	c, err := hex.DecodeString("6c73d5240a948c86981bc294814d")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PT: %b\n", m)
	fmt.Printf("CT: %b\n", c)

	key := xorBytes(m, c)
	fmt.Println(key)

	fmt.Println(string(xorBytes(c, key)))
	fmt.Println(hex.EncodeToString(xorBytes(m, key)))
	fmt.Println(hex.EncodeToString(xorBytes(m2, key)))
}

func xorBytes(b1, b2 []byte) []byte {
	if len(b1) != len(b2) {
		panic("length mismatch")
	}
	xor := make([]byte, len(b1))

	for i := range b1 {
			xor[i] = b1[i] ^ b2[i]
	}

	return xor
}
