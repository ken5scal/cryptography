package main

import (
	"fmt"
	"strconv"
)

func main() {
	src := []byte("hogehogehogehoge")
	s0 := uint32(src[0])<<24 | uint32(src[1])<<16 | uint32(src[2])<<8 | uint32(src[3])
	fmt.Println(src[3], fmt.Sprintf("%b", src[3]), fmt.Sprintf("%b", uint32(src[3])))
	fmt.Println(src[2], fmt.Sprintf("%b", src[2]), fmt.Sprintf("%b", uint32(src[2])<<8))
	fmt.Println(src[1], fmt.Sprintf("%b", src[1]), fmt.Sprintf("%b", uint32(src[1])<<16))
	fmt.Println(src[0], fmt.Sprintf("%b", src[0]), fmt.Sprintf("%b", uint32(src[0])<<24))
	fmt.Println(fmt.Sprintf("%b", s0))
	fmt.Println([]rune("e"))
	fmt.Println(strconv.QuoteRune(101))

	ff := uint32(0xff) // 00000000000000000000000011111111
	fmt.Println(fmt.Sprintf("%b", ff))
	fmt.Println(fmt.Sprintf("%b", s0>>24), fmt.Sprintf("%b", s0>>24&ff))
	fmt.Println(fmt.Sprintf("%b", s0>>16), fmt.Sprintf("%b", s0>>16&ff))
}
