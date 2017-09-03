package aes

import (
	"testing"
)

func int32tobyte(b []byte, v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[3] = byte(v)
	b[2] = byte(v >> 8)
	b[1] = byte(v >> 16)
	b[0] = byte(v >> 24)
}

func TestGenerateStateBlock(t *testing.T) {
	input := "hogefugaabcdefgh"
	src := []byte(input)
	state, err := GenerateStateBlock(src)

	if err != nil {
		t.Error(err.Error())
	}

	buf := make([]byte, len(state)*4)
	for i, v := range state {
		int32tobyte(buf[i*4:], v)
	}

	for i := range src {
		if buf[i] != src[i] {
			t.Errorf("GenerateStateBlock(%b) = %b, want %b", src, state[i], src[i])
		}
	}
}

//func TestSubBytes(t *testing.T) {
//	input := "hogefugaabcdefgh"
//	state, _ := GenerateStateBlock([]byte(input))
//	newState, err := SubBytes(state)
//
//	if err != nil {
//		t.Error(err.Error())
//	}
//	buf := make([]byte, len(state)*4)
//	for i, v := range state {
//		int32tobyte(buf[i*4:], v)
//	}
//
//	fmt.Printf("%b\n", buf)
//	fmt.Printf("%b\n", newState)
//	t.Error("No test has been written.") // I don't think you can...
//}
