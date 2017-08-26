package main

import (
	"errors"
	"math/big"
	"math/rand"
	"time"
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
	if r.Cmp(bigTwo) == -1 || t.Cmp(bigOne) == -1 {
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
	if r.Cmp(bigTwo) == 0 {
		return true, nil
	}

	if r.Cmp(bigThree) == -1 || t.Cmp(bigOne) == -1 {
		return false, errors.New("r must be larger than or equal to 3, and t must be larger than or equal to 1")
	}

	s, k := findSandK(r) // 2^s*k = r - 1
	for i := 0; i < int(t.Int64()); i++ {
		if !millerTest(r, k, s) {
			return false, nil
		}
	}

	return true, nil
}

func millerTest(r *big.Int, k *big.Int, s *big.Int) bool {
	rMinusOne := new(big.Int).Sub(r, bigOne)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	a := new(big.Int).Rand(rnd, new(big.Int).Sub(r, big.NewInt(4)))
	a.Add(a, bigTwo)

	x := new(big.Int).Exp(a, k, r)

	if x.Cmp(bigOne) == 0 || x.Cmp(rMinusOne) == 0 {
		return true
	}

	for n := 0; n < int(s.Int64()); n++ {
		x.Exp(x, bigTwo, r)
		if x.Cmp(bigOne) == 0 {
			return false
		} else if x.Cmp(rMinusOne) == 0 {
			return true
		}
	}

	return false
}

// findSandK finds s and k which satisfies r - 1 = 2^s * k
func findSandK(r *big.Int) (*big.Int, *big.Int) {
	r_cp := big.NewInt(r.Int64())
	s := new(big.Int)
	k := new(big.Int).Sub(r, bigOne)
	for r_cp.Mod(k, bigTwo).Cmp(bigZero) == 0 {
		s.Add(s, bigOne)
		k.Div(k, bigTwo)
	}
	return s, k
}
