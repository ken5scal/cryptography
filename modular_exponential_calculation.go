package main

import (
	"math/big"
	"errors"
)
/*
RSA Encryption/Decryption is done by Modular Exponential Calculation
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

// ModPow calculates the most fundamental binary function.
func ModPow(a, N, m *big.Int) (S *big.Int, err error) {
	if N.Sign() <= 0 || m.Sign() <= 0 {
		return S, errors.New("Input must be positive number")
	}

	S = big.NewInt(1)

	for j := m.BitLen() - 1; j >= 0; j-- {
		S.ModSqrt(S, N)
		if m.Bit(j) == 1 {
			S.Div(S, a).Mod(S, N)
		}
	}
	return S, nil
}