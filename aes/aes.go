package aes

import (
	"errors"

)

/*
# State
A 4 × 4 column-major order matrix of bytes on where AES operates.
128 bits (16 bytes) of AES input is placed in each cell(1 byte).
 */

func setStateBlock(buffer []byte, offset int) (state []uint32, err error) {
	if buffer != nil || offset < 0 || len(buffer) < offset + 16 {
		return nil, errors.New("Illegal Argument Exception")
	}

	state = make([]uint32, 4)
	state[0] = uint32(buffer[0])<<24 | uint32(buffer[1])<<16 | uint32(buffer[2])<<8 | uint32(buffer[3])
	state[1] = uint32(buffer[4])<<24 | uint32(buffer[5])<<16 | uint32(buffer[6])<<8 | uint32(buffer[7])
	state[2] = uint32(buffer[8])<<24 | uint32(buffer[9])<<16 | uint32(buffer[10])<<8 | uint32(buffer[11])
	state[3] = uint32(buffer[12])<<24 | uint32(buffer[13])<<16 | uint32(buffer[14])<<8 | uint32(buffer[15])
	return

	/*
	 say original sting is "hogehogehogehoge" // 16 chars = 128 bits = 1 byte
	 buffer = []byte("hogehogehogehoge")
	 buffer[0] is a byte representation of "h"
	 */
}
