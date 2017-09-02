package aes

import (
	"errors"
	"fmt"
)

/*
# State
A 4 × 4 column-major order matrix of bytes on where AES operates.
128 bits (16 bytes) of AES input is placed in each cell(1 byte).
 */

func setStateBlock(buffer []byte, offset int) error {
	if buffer != nil || offset < 0 || len(buffer) < offset + 16 {
		return errors.New("Illegal Argument Exception")
	}
	return nil
}

func getStateBlock(buffer []byte, offset int) {

}
