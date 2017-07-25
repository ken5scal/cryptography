package main

import (
	"math/big"
	"fmt"
	"errors"
)

func main() {
	g, _ := Gcd(big.NewInt(2), big.NewInt(10))
	fmt.Println(g)
}


func Gcd(a, b *big.Int) (*big.Int, error) {
	if a.Sign() <= 0 || b.Sign() <= 0 {
		return nil, errors.New("Input must be positive number")
	}

	g := big.NewInt(1)
	if a.Cmp(b) < 0 {
		tmp := a; a = b; b= tmp
	}

	n := big.NewInt(2)
	var tmpA big.Int
	var tmpB big.Int

	for n.Cmp(b) <= 0 {
		for  {
			if tmpA.Mod(a, n).Cmp(big.NewInt(0)) == 0 &&
				tmpB.Mod(b, n).Cmp(big.NewInt(0)) == 0 {
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
