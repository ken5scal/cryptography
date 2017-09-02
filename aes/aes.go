package aes

import (
	"errors"

)

/*
# State
A 4 × 4 column-major order matrix of bytes on where AES operates.
128 bits (16 bytes) of AES input is placed in each cell(1 byte).
 */

// setStateBlock generates block(state) to be calculated by AES.
// currently ignoring `offset` usage.
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
	 Each column of a state matrix sores 4 consecutive buffers.
	 For example first column stores, buffer[0], buffer[1], buffer[2], buffer[3].
	 In above, uint32 is prepared for a column, and each buffer is stored consecutively.
	 For example, buffer[0] is stored in first 8 bits of uint32, while buffer[1] is in second 8 bits.

	 EX:
	 say original sting is "hogehogehogehoge" // 16 chars = 128 bits = 1 byte
	 buffer = []byte("hogehogehogehoge")
	 buffer[0] is a byte representation of "h", which is 104 in RUNE (or Code point).
	 rune 104 can be represented as 01101000 in binary, which is 8 bits long.
	 so to store that in state[0],  01101000 "h" is left shifted by 24.
	 do the same thing for 01101111 "o", 01100111 "g", but left shift by 16 and 8.
	 finally store 01100101 "e" in last 8 bits of state[0].

	 Do the same thing for buffer[4:8] in state[1], buffer[8:12] in state[2], buffer[12:16] in state[3]
	 */
}
