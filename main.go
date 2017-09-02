package main

import (
	//"bytes"
	"fmt"
	"strconv"
)

func main() {
	ff := "hogehogehogehoge" // 16 chars = 128 bits = 1 byte
	src := []byte(ff)
	fmt.Println(len(src))
	s0 := uint32(src[0])<<24 | uint32(src[1])<<16 | uint32(src[2])<<8 | uint32(src[3])
	fmt.Println(src[3], uint32(src[3]), fmt.Sprintf("%b",uint32(src[3])))
	fmt.Println(uint32(src[2])<<8, fmt.Sprintf("%b",uint32(src[2])<<8))
	fmt.Println(uint32(src[1])<<16, fmt.Sprintf("%b",uint32(src[1])<<16))
	fmt.Println(uint32(src[0])<<24, fmt.Sprintf("%b",uint32(src[0])<<24))
	fmt.Println(fmt.Sprintf("%b", s0))
	fmt.Println([]rune("e"))
	fmt.Println(strconv.QuoteRune(101))
}
