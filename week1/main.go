package main

import (
	"fmt"
	"encoding/hex"
	"log"
)

func main() {
	m := []byte("attack at dawn")

	ct := "6c73d5240a948c86981bc294814d" //hex
	src := []byte(ct)
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PT: %b\n", m)
	fmt.Printf("CT: %b\n", dst)


	var key []byte
	for i, v := range m {
		key = append(key, v ^ dst[i])
	}
	fmt.Printf("%b\n", key)
	for i, v := range key {
		fmt.Println(v ^ dst[i])
	}
}
