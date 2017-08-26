package main

import (
	"errors"
	"math/big"
	"math/rand"
)

var bigZero = big.NewInt(0)
var bigOne = big.NewInt(1)
var bigTwo = big.NewInt(2)
var bigThree = big.NewInt(3)

// IsPrimeByFermatTest judges whether r is prime number using Probabilistic Prime Test algorithm, Fermat
// t is usually 1.
// Down side is that the test does not quantify a probability of r being prime number.
// It just repeats t times to check {a^(p-1) mod p} is 1  (a ^(p-1) - 1 is multiple of p)
func IsPrimeByFermatTest(r *big.Int, t *big.Int) (bool, error) {

	if r.Cmp(bigThree) == -1 || t.Cmp(bigOne) == -1 {
		return false, errors.New("r must be larger than or equal to 3, and t must be larger than or equal to 1")
	}

	S := new(big.Int)

	for j := 1; j < int(t.Int64()); j++ {
		// Pre-assumption is that a is universally random.
		a := big.NewInt(rand.Int63n(r.Int64()) - 1)
		S.Exp(a, S.Sub(r, bigOne), r)
		if S.Cmp(bigOne) != 0 {
			return false, nil
		}
	}
	return true, nil
}

// IsPrimeByMillerRabinTest is more efficient thant Fermat test.
func IsPrimeByMillerRabinTest(r *big.Int, t *big.Int) (bool, error) {
	if r.Cmp(bigThree) == -1 || t.Cmp(bigOne) == -1 {
		return false, errors.New("r must be larger than or equal to 3, and t must be larger than or equal to 1")
	}

	// 2^s*K = r -r
	s := 0
	d := new(big.Int).Sub(r, bigOne)
	for d.Mod(d, bigTwo).Cmp(bigZero) != 0{
		s++
		d.Div(d, bigTwo)
	}

	return true, nil
}

func findSAndK(r  *big.Int) (s, k *big.Int) {
	d := new(big.Int).Sub(r, bigOne)
	for d.Mod(k, bigTwo).Cmp(bigZero) == 0 {
		s.Add(s, bigOne)
		k.Div(k, bigTwo)
	}
	return
}