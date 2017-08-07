package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

/*
RSA Encryption/Decryption is done by Modular Exponential Calculation(MEC)
S = a^m mod N
 Encryption:
  - S: Cipher Text
  - a: Plain Text
  - m: Public Exponent公開指数(e)
  - N: Public Modulus(N)
 Decryption:
  - S: Plain Text
  - a: Cipher Text
  - m: Secret Exponent秘密指数(d)
  - N: Public Modulus(N)
*/

func getWindow(m *big.Int, w int) (mjw int64) {
	for j := (m.BitLen()+w-1)/w - 1; j >= 0; j-- {
		mjw = int64(0)
		for i := w - 1; i >= 0; i-- {
			mjw <<= 1
			if m.Bit(j*w+i) != 0 {
				mjw |= 1
			}
		}
		fmt.Println(mjw)
	}
	return mjw
}

// ModPow2wary is another method to do MEC using Window(2w-ary) ModPow
func ModPow2wary(a, m, N *big.Int, w int) (*big.Int, error) {
	if N.Sign() <= 0 || m.Sign() <= 0 {
		return nil, errors.New("Input must be positive number")
	}

	length := int64(math.Pow(2, float64(w)))
	table := make([]*big.Int, length)
	table[0] = big.NewInt(1)
	for k := 1; k < int(length); k++ {
		table[k] = big.NewInt(table[k-1].Int64() * a.Int64() % N.Int64())
	}

	S := big.NewInt(1)
	for j := (m.BitLen()+w-1)/w - 1; j >= 0; j-- {
		for i := 0; i < w; i++ {
			S.Mul(S, S).Mod(S, N)
		}

		mjw := int64(0)
		for i := w - 1; i >= 0; i-- {
			mjw <<= 1
			if m.Bit(j*w+i) != 0 {
				mjw |= 1
			}
		}
		S.Mul(S, table[mjw]).Mod(S, N)
	}

	return S, nil
}

// ModPow calculates the most fundamental binary function.
func ModPow(a, m, N *big.Int) (*big.Int, error) {
	if N.Sign() <= 0 || m.Sign() <= 0 {
		return nil, errors.New("Input must be positive number")
	}

	S := big.NewInt(1)

	for j := m.BitLen() - 1; j >= 0; j-- {
		S.Mul(S, S).Mod(S, N)

		if m.Bit(j) == 1 {
			S.Mul(S, a).Mod(S, N)
		}
	}
	return S, nil
}
