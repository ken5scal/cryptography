package main

import (
	"errors"
	"math/big"
	"math/rand"
)

var bigOne = big.NewInt(1)
var bigThree = big.NewInt(3)

func IsPrimeByFermat(r *big.Int, t *big.Int) (bool, error) {

	if r.Cmp(bigThree) == -1 || t.Cmp(bigOne) == -1 {
		return false, errors.New("r must be larger than or equal to 3, and t must be larger than or equal to 1")
	}

	a := big.NewInt(rand.Int63n(r.Int64()) - 1)
	S := new(big.Int)

	for j := 1; j < int(t.Int64()); j++ {
		S.Exp(a, S.Sub(r, bigOne), r)
		if S.Cmp(bigOne) != 0 {
			return false, nil
		}
	}
	return true, nil
}
