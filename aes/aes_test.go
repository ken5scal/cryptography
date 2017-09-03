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
//	testState := []uint32{
//		01101000011011110110011101100101, //hoge
//		01101000011011110110011101100101, //hoge
//		01101000011011110110011101100101, //hoge
//		01101000011011110110011101100101} //hoge
//	newState, err := SubBytes(testState)
//
//	expectedState := []uint32{11010101, 11010101, 11010101, 11010101}
//
//	if err != nil {
//		t.Errorf("SubBytes[%#x] = %#x, want %#x", testState, newState, expectedState)
//	}
//
//	for i := range expectedState {
//		if testState[i] != expectedState[i] {
//			t.Errorf("SubBytes[%#x] = %#x, want %#x", testState, newState, expectedState)
//		}
//	}
//}
