package aes

import (
	"testing"
	"encoding/binary"
)

func TestGenerateStateBlock(t *testing.T) {
	input := "hogefugaabcdefgh"
	src := []byte(input)
	state, err := GenerateStateBlock(src)
	if err != nil {
		t.Error(err.Error())
	}

	buf := make([]byte, len(state)*4)

	for i, v := range state {
		binary.LittleEndian.PutUint32(buf[i*4:], v)
	}

	for i := range src {
		if buf[i] != src[i] {
			t.Errorf("GenerateStateBlock(%b) = %b, want %b", src, state[i], src[i])
		}
	}
}

func TestSubBytes(t *testing.T) {
	input := "hogefugaabcdefgh"
	state, _ := GenerateStateBlock([]byte(input))
	newState, err := SubBytes(state)

	if err != nil {
		t.Error(err.Error())
	}

	revState, err := InvSubBytes(newState)

	if err != nil {
		t.Error(err.Error())
	}

	for i, v := range state {
		if v != revState[i] {
			t.Errorf("InvSubBytes(SubBytes(%b)) = %b, want %b", state, revState, state)
		}
	}
}
