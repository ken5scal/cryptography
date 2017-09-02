package main

import (
	//"bytes"
	"fmt"
)

func main() {
	ff := "hogehogehogehoge" // 16 chars = 128 bits = 1 byte

	//var buffer bytes.Buffer
	//for i := 0; i < len(ff);  i++ {
	//	fmt.Fprintf(&buffer, "%b", ff[i])
	//}

	src := []byte(ff)

	//var s0, s1, s2, s3 uint32

	s0 := uint32(src[0])<<24 | uint32(src[1])<<16 | uint32(src[2])<<8 | uint32(src[3])
	//s1 = uint32(src[4])<<24 | uint32(src[5])<<16 | uint32(src[6])<<8 | uint32(src[7])
	//s2 = uint32(src[8])<<24 | uint32(src[9])<<16 | uint32(src[10])<<8 | uint32(src[11])
	//s3 = uint32(src[12])<<24 | uint32(src[13])<<16 | uint32(src[14])<<8 | uint32(src[15])

	fmt.Println(uint32(src[3]), fmt.Sprintf("%b",uint32(src[3])))
	fmt.Println(uint32(src[2])<<8, fmt.Sprintf("%b",uint32(src[2])<<8))
	fmt.Println(uint32(src[1])<<16, fmt.Sprintf("%b",uint32(src[1])<<16))
	fmt.Println(uint32(src[0])<<24, fmt.Sprintf("%b",uint32(src[0])<<24))
	fmt.Println(fmt.Sprintf("%b", uint32(src[0])<<24 | uint32(src[1])<<16 | uint32(src[2])<<8 | uint32(src[3])))
	fmt.Println(s0)
}
