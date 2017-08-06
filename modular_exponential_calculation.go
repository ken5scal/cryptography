package main

import (
	"math/big"
	"errors"
)

func ModPow(a, N, m big.Int) (S *big.Int, err error) {
	if N.Sign() <= 0 || m.Sign() <= 0 {
		return S, errors.New("Input must be positive number")
	}

	S = big.NewInt(1)

	for j := m.BitLen() - 1; j >= 0; j-- {
		S.ModSqrt(S, &N)
		if m.Bit(j) == 1 {
			S.Div(S, &a).Mod(S, &N)
		}
	}
	return S, nil
}