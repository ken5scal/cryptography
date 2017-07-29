package main

import (
	"math/big"
	"fmt"
	"errors"
)

func main() {
	a := big.NewInt(780)
	b := big.NewInt(600)
	g, _ := Gcd(a, b)
	fmt.Println(g)

	a = big.NewInt(780)
	b = big.NewInt(600)
	g = EuclidGCD(a, b)
	fmt.Println(g)
}

func EuclidGCD(aOrig, bOrig *big.Int) *big.Int {
	// Making sure immutability
	a := big.NewInt(aOrig.Int64())
	b := big.NewInt(bOrig.Int64())
	// Nonzero Initialization
	for r := big.NewInt(1);; {
		r.Mod(a, b)
		if r.Cmp(big.NewInt(0)) == 0 {
			return b
		}
		a.Set(b)
		b.Set(r)
	}
	return b
}

// 試行 割り算
func Gcd(a, b *big.Int) (*big.Int, error) {
	if a.Sign() <= 0 || b.Sign() <= 0 {
		return nil, errors.New("Input must be positive number")
	}

	var tmp big.Int
	g := big.NewInt(1)
	if a.Cmp(b) < 0 {
		tmp := a
		a = b
		b = tmp
	}

	n := big.NewInt(2)
	for n.Cmp(b) <= 0 {
		for  {
			if tmp.Mod(a, n).Cmp(big.NewInt(0)) == 0 &&
				tmp.Mod(b, n).Cmp(big.NewInt(0)) == 0 {
				g.Mul(g, n)
				a.Div(a, n)
				b.Div(b, n)
			} else {
				break
			}
		}
		n.Add(n, big.NewInt(1))
	}

	return g, nil
}
