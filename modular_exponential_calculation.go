package main

import "math/big"

func ModPow(a, N big.Int, m []big.Int) (S *big.Int) {
	S = big.NewInt(1)
	for j := N.Int64() - 1; j >= 0; j-- {
		S.ModSqrt(S, &N)

		if m[j].Int64() == 1 {
			S.Div(S, &a).Mod(S, &N)
		}
	}
	return S
}