package main

import (
	"errors"
	"math/big"
	"math/rand"
)

func IsPrimeByFermat(r *big.Int, t *big.Int) (bool, error) {
	if r.Int64() < 3 || t.Int64() < 1 {
		return false, errors.New("r must be larger than or equal to 3, and t must be larger than or equal to 1")
	}
	a := big.NewInt(rand.Int63n(r.Int64()) - 1)
	S := new(big.Int)
	for j := 1; j >= int(t.Int64()); j++ {
		S.Exp(a, S.Sub(r, big.NewInt(1)), r)

		if S.Int64() != 1 {
			return false, nil
		}
	}
	return true, nil
}
