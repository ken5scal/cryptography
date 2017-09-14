package main

import (
	"fmt"
	"cryptography/aes"
)

func main() {
	input := "hogehogehogehoge"
	src := []byte(input)
	s0 := uint32(src[0])<<24 | uint32(src[1])<<16 | uint32(src[2])<<8 | uint32(src[3])
	fmt.Println(string(src[3]), src[3], fmt.Sprintf("%b", src[3]), fmt.Sprintf("%b", uint32(src[3])))
	fmt.Println(string(src[2]), src[2], fmt.Sprintf("%b", src[2]), fmt.Sprintf("%b", uint32(src[2])<<8))
	fmt.Println(string(src[1]), src[1], fmt.Sprintf("%b", src[1]), fmt.Sprintf("%b", uint32(src[1])<<16))
	fmt.Println(string(src[0]), src[0], fmt.Sprintf("%b", src[0]), fmt.Sprintf("%b", uint32(src[0])<<24))
	fmt.Println(fmt.Sprintf("%b", s0))
	fmt.Println([]rune("e"))

	state, _ := aes.GenerateStateBlock(src)
	fmt.Println(fmt.Sprintf("%b", state))

	ff := uint32(0xff) // 00000000000000000000000011111111
	fmt.Println(fmt.Sprintf("%b", ff))
	fmt.Println(fmt.Sprintf("%b", s0>>24), fmt.Sprintf("%b", s0>>24&ff))
	fmt.Println(fmt.Sprintf("%b", s0>>16), fmt.Sprintf("%b", s0>>16&ff))
}
