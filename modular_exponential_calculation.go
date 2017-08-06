package main

import (
	"errors"
	"math/big"
	"fmt"
	"strconv"
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

func GetWindow(m *big.Int, w int) {
	for j := (m.BitLen() + w -1)/w - 1; j >= 0; j-- {
		mjw := int64(0)
		for i := w - 1; i >= 0 ; i-- {
			fmt.Printf("Before <<=: %v\n", strconv.FormatInt(mjw, 2))
			mjw <<= 1
			fmt.Printf("After <<=: %v\n", strconv.FormatInt(mjw, 2))
			if m.Bit(j * w + i) != 0 {
				fmt.Printf("Before |=: %v\n", strconv.FormatInt(mjw, 2))
				mjw |= 1
				fmt.Printf("After |=: %v\n", strconv.FormatInt(mjw, 2))
			}
		}
		fmt.Println(mjw)
	}
}

// ModPow2wary is another method to do MEC using Window(2w-ary) ModPow
func ModPow2wary(a, m, N *big.Int, w int) (*big.Int, error) {
	if N.Sign() <= 0 || m.Sign() <= 0 {
		return nil, errors.New("Input must be positive number")
	}

	table := make([]*big.Int, 2 ^ w)
	table[0] = big.NewInt(0)
	for k := 1; k < 2 ^ w; k++ {
		table[k].Mul(table[k-1], a).Mod(table[k], N)
	}

	S := big.NewInt(1)
	for j := m.BitLen()/w - 1; j >= 0; j-- {
		for i := 0; i < w; i++ {
			S.Mul(S, S).Mod(S, N)
		}

		for i := w - 1; i >=0; i-- {
			if m.Bit(j) != 0 {
				//S.Mul(S, a).Mul(S, hoge).Mod(S, N)
			}
		}
	}

	return nil, nil
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
