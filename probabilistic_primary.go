package main

import (
	"errors"
	"math/big"
	"math/rand"
	"fmt"
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

	// 2^s*k = r - 1
	s, k := findSandK(r)
	r_minus_1 := new(big.Int).Sub(r, bigOne)
	fmt.Println(s, k, r_minus_1)

	for i := 0; i < int(t.Int64()); i++ {
		a := big.NewInt(rand.Int63n(r.Int64()) + 2)
		x := new(big.Int).Exp(a, k, r)
		if x.Cmp(bigOne) == 0 || x.Cmp(r_minus_1) == 0 {
			continue
		}

		n := 1
		for ; n < int(s.Int64()); n++ {
			x.Exp(x, bigTwo, r)
			if x.Cmp(bigOne) == 0 {
				return false, nil
			} else if x.Cmp(r_minus_1) == 0 {
				break
			}
		}
		if s.Int64() == int64(n) {
			return false, nil
		}
	}

	return true, nil
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
