package aes

import (
	"errors"
)

/*
# State
A 4 × 4 column-major order matrix of bytes on where AES operates.
128 bits (16 bytes) of AES input is placed in each cell(1 byte).
*/

// GenerateStateBlock generates block(state) to be calculated by AES.
// currently ignoring `offset` usage.
//func GenerateStateBlock(buffer []byte, offset int) (state []uint32, err error) {
func GenerateStateBlock(buffer []byte) (state []uint32, err error) {
	//if buffer != nil || offset < 0 || len(buffer) < offset+16 {
	if buffer == nil {
		return nil, errors.New("Illegal Argument Exception")
	}

	state = make([]uint32, 4)
	//state[0] = uint32(buffer[0])<<24 | uint32(buffer[1])<<16 | uint32(buffer[2])<<8 | uint32(buffer[3])
	//state[1] = uint32(buffer[4])<<24 | uint32(buffer[5])<<16 | uint32(buffer[6])<<8 | uint32(buffer[7])
	//state[2] = uint32(buffer[8])<<24 | uint32(buffer[9])<<16 | uint32(buffer[10])<<8 | uint32(buffer[11])
	//state[3] = uint32(buffer[12])<<24 | uint32(buffer[13])<<16 | uint32(buffer[14])<<8 | uint32(buffer[15])
	//state[0] = uint32(buffer[0]) | uint32(buffer[1])<<8 | uint32(buffer[2])<<16 | uint32(buffer[3])<<24
	//state[1] = uint32(buffer[4]) | uint32(buffer[5])<<8 | uint32(buffer[6])<<16 | uint32(buffer[7])<<24
	//state[2] = uint32(buffer[8]) | uint32(buffer[9])<<8 | uint32(buffer[10])<<16 | uint32(buffer[11])<<24
	//state[3] =  uint32(buffer[12]) | uint32(buffer[13])<<8 | uint32(buffer[14])<<16 | uint32(buffer[15])<<24
	for i := 0; i < len(state); i++ {
		state[i] = uint32(buffer[i * 4]) |
			uint32(buffer[i * 4 + 1]) << 8 |
			uint32(buffer[i * 4 + 2]) << 16 |
			uint32(buffer[i * 4 + 3]) << 24
	}
	return

	/*
	 Each column of a state matrix sores 4 consecutive buffers.
	 For example first column stores, buffer[0], buffer[1], buffer[2], buffer[3].
	 In above, uint32 is prepared for a column, and each buffer is stored consecutively.
	 For example, buffer[0] is stored in first 8 bits of uint32, while buffer[1] is in second 8 bits.

	 EX:
	 say original sting is "hogehogehogehoge" // 16 chars  = 16 bytes = 128 bits
	 buffer = []byte("hogehogehogehoge")
	 buffer[0] is a byte representation of "h", which is 104 in RUNE (or Code point).
	 rune 104 can be represented as 01101000 in binary, which is 8 bits long.
	 so to store that in state[0],  01101000 "h" is left shifted by 24.
	 do the same thing for 01101111 "o", 01100111 "g", but left shift by 16 and 8.
	 finally store 01100101 "e" in last 8 bits of state[0].

	 Do the same thing for buffer[4:8] in state[1], buffer[8:12] in state[2], buffer[12:16] in state[3]
	*/
}

// AddRoundKey XOR round key RK_i to a state block
func AddRoundKey(state, rk []uint32) (newState []uint32, err error) {
	if len(state) != 4 || len(rk) != 4 || state == nil || rk == nil {
		return nil, errors.New("Illegal Argument Exception")
	}
	newState = make([]uint32, 4)
	newState[0] = state[0] ^ rk[0]
	newState[1] = state[1] ^ rk[1]
	newState[2] = state[2] ^ rk[2]
	newState[3] = state[3] ^ rk[3]

	return
}

// SubBytes applies nonlinear transformation per state.
// There is so called a SBox (substitution-box), which is a table that converts 1 byte to new byte.
// One byte is split into two 4bits. First 4 bits corresponds to x, and last 4 bits corresponds to y.
// A cell at (x, y) in SBox will be the substituted result.
// For example, a byte 10110101 is split into 1011(b), 0101(5).
// Reading (b, 5) in a SBox, the output will be 'd5'.
func SubBytes(state []uint32) (newState []uint32, err error) {
	if len(state) != 4 || state == nil {
		return nil, errors.New("Illegal Argument Exception")
	}

	ff := uint32(0xff) // 00000000000000000000000011111111
	newState = make([]uint32, 4)

	for i := 0; i < len(newState); i++ {
		newState[i] = uint32(sbox0[state[i]>>24])<<24 |
			uint32(sbox0[state[i]>>16&ff])<<16 | // &ff eliminates all bits except last 8 bits.
			uint32(sbox0[state[i]>>8&ff])<<8 |
			uint32(sbox0[state[i]&ff])
	}

	return
}

// InvSubBytes reverses Substituted state back to original state.
func InvSubBytes(state []uint32) (newState []uint32, err error) {
	if len(state) != 4 || state == nil {
		return nil, errors.New("Illegal Argument Exception")
	}

	ff := uint32(0xff) // 00000000000000000000000011111111
	newState = make([]uint32, 4)

	for i := 0; i < len(newState); i++ {
		newState[i] = uint32(sbox1[state[i]>>24])<<24 |
			uint32(sbox1[state[i]>>16&ff])<<16 | // &ff eliminates all bits except last 8 bits.
			uint32(sbox1[state[i]>>8&ff])<<8 |
			uint32(sbox1[state[i]&ff])
	}

	return
}

// ShiftRows shifts each row by {row index} byte.
// For example, 1st row is shifted by 0 byte(so sty the same), 2nd row is shifted by 1 byte, 3rd row by 2 bytes, and 4th row by 3 bytes
func ShiftRows(state []uint32) (newState []uint32) {
	ff := uint32(0xff) // 00000000000000000000000011111111
	newState = make([]uint32, 4)

	s00 := state[0]>>24
	s10 := state[1]>>16&ff
	s20 := state[2]>>8&ff
	s30 := state[3]&ff

	newState[0] = uint32(s00)<<24 | uint32(s10)<<16 | uint32(s20)<<8 |uint32(s30)

	s01 := state[1]>>24
	s11 := state[2]>>16&ff
	s21 := state[3]>>8&ff
	s31 := state[0]&ff

	newState[1] = uint32(s01)<<24 | uint32(s11)<<16 | uint32(s21)<<8 |uint32(s31)

	s02 := state[2]>>24
	s12 := state[3]>>16&ff
	s22 := state[0]>>8&ff
	s32 := state[1]&ff

	newState[2] = uint32(s02)<<24 | uint32(s12)<<16 | uint32(s22)<<8 |uint32(s32)

	s03 := state[3]>>24
	s13 := state[0]>>16&ff
	s23 := state[1]>>8&ff
	s33 := state[2]&ff

	newState[3] = uint32(s03)<<24 | uint32(s13)<<16 | uint32(s23)<<8 |uint32(s33)

	return
}

// InvShiftRows does ...
func InvShiftRows(state []uint32) (newState []uint32) {
	ff := uint32(0xff) // 00000000000000000000000011111111
	newState = make([]uint32, 4)

	s00 := state[0]>>24
	s10 := state[1]>>16&ff
	s20 := state[2]>>8&ff
	s30 := state[3]&ff

	newState[0] = uint32(s00)<<24 | uint32(s10)<<16 | uint32(s20)<<8 |uint32(s30)

	s01 := state[1]>>24
	s11 := state[2]>>16&ff
	s21 := state[3]>>8&ff
	s31 := state[0]&ff

	newState[1] = uint32(s01)<<24 | uint32(s11)<<16 | uint32(s21)<<8 |uint32(s31)

	s02 := state[2]>>24
	s12 := state[3]>>16&ff
	s22 := state[0]>>8&ff
	s32 := state[1]&ff

	newState[2] = uint32(s02)<<24 | uint32(s12)<<16 | uint32(s22)<<8 |uint32(s32)

	s03 := state[3]>>24
	s13 := state[0]>>16&ff
	s23 := state[1]>>8&ff
	s33 := state[2]&ff

	newState[3] = uint32(s03)<<24 | uint32(s13)<<16 | uint32(s23)<<8 |uint32(s33)

	return
}

// FIPS-197 Figure 7. S-box substitution values in hexadecimal format.
// Copied from go/src/crypto/aes/const.go
var sbox0 = [256]byte{
	0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76,
	0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0,
	0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15,
	0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75,
	0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84,
	0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf,
	0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8,
	0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2,
	0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73,
	0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb,
	0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79,
	0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08,
	0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a,
	0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e,
	0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf,
	0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16,
}

// FIPS-197 Figure 14.  Inverse S-box substitution values in hexadecimal format.
// Use this in decryption.
var sbox1 = [256]byte{
	0x52, 0x09, 0x6a, 0xd5, 0x30, 0x36, 0xa5, 0x38, 0xbf, 0x40, 0xa3, 0x9e, 0x81, 0xf3, 0xd7, 0xfb,
	0x7c, 0xe3, 0x39, 0x82, 0x9b, 0x2f, 0xff, 0x87, 0x34, 0x8e, 0x43, 0x44, 0xc4, 0xde, 0xe9, 0xcb,
	0x54, 0x7b, 0x94, 0x32, 0xa6, 0xc2, 0x23, 0x3d, 0xee, 0x4c, 0x95, 0x0b, 0x42, 0xfa, 0xc3, 0x4e,
	0x08, 0x2e, 0xa1, 0x66, 0x28, 0xd9, 0x24, 0xb2, 0x76, 0x5b, 0xa2, 0x49, 0x6d, 0x8b, 0xd1, 0x25,
	0x72, 0xf8, 0xf6, 0x64, 0x86, 0x68, 0x98, 0x16, 0xd4, 0xa4, 0x5c, 0xcc, 0x5d, 0x65, 0xb6, 0x92,
	0x6c, 0x70, 0x48, 0x50, 0xfd, 0xed, 0xb9, 0xda, 0x5e, 0x15, 0x46, 0x57, 0xa7, 0x8d, 0x9d, 0x84,
	0x90, 0xd8, 0xab, 0x00, 0x8c, 0xbc, 0xd3, 0x0a, 0xf7, 0xe4, 0x58, 0x05, 0xb8, 0xb3, 0x45, 0x06,
	0xd0, 0x2c, 0x1e, 0x8f, 0xca, 0x3f, 0x0f, 0x02, 0xc1, 0xaf, 0xbd, 0x03, 0x01, 0x13, 0x8a, 0x6b,
	0x3a, 0x91, 0x11, 0x41, 0x4f, 0x67, 0xdc, 0xea, 0x97, 0xf2, 0xcf, 0xce, 0xf0, 0xb4, 0xe6, 0x73,
	0x96, 0xac, 0x74, 0x22, 0xe7, 0xad, 0x35, 0x85, 0xe2, 0xf9, 0x37, 0xe8, 0x1c, 0x75, 0xdf, 0x6e,
	0x47, 0xf1, 0x1a, 0x71, 0x1d, 0x29, 0xc5, 0x89, 0x6f, 0xb7, 0x62, 0x0e, 0xaa, 0x18, 0xbe, 0x1b,
	0xfc, 0x56, 0x3e, 0x4b, 0xc6, 0xd2, 0x79, 0x20, 0x9a, 0xdb, 0xc0, 0xfe, 0x78, 0xcd, 0x5a, 0xf4,
	0x1f, 0xdd, 0xa8, 0x33, 0x88, 0x07, 0xc7, 0x31, 0xb1, 0x12, 0x10, 0x59, 0x27, 0x80, 0xec, 0x5f,
	0x60, 0x51, 0x7f, 0xa9, 0x19, 0xb5, 0x4a, 0x0d, 0x2d, 0xe5, 0x7a, 0x9f, 0x93, 0xc9, 0x9c, 0xef,
	0xa0, 0xe0, 0x3b, 0x4d, 0xae, 0x2a, 0xf5, 0xb0, 0xc8, 0xeb, 0xbb, 0x3c, 0x83, 0x53, 0x99, 0x61,
	0x17, 0x2b, 0x04, 0x7e, 0xba, 0x77, 0xd6, 0x26, 0xe1, 0x69, 0x14, 0x63, 0x55, 0x21, 0x0c, 0x7d,
}
