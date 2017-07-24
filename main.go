package main

import (
	"math/big"
	"fmt"
)

func main() {
	g := Gcd(big.NewInt(780), big.NewInt(660))
	fmt.Println(g)
}


func Gcd(a, b *big.Int) *big.Int {
	g := big.NewInt(1)
	if a.Cmp(b) < 0 {
		tmp := a; a = b; b= tmp
	}

	n := big.NewInt(2)
	for n.Cmp(b) < 0 {
		for  {
			var tmpA big.Int
			var tmpB big.Int

			if tmpA.Mod(a, n).Cmp(big.NewInt(0)) == 0 &&
				tmpB.Mod(b, n).Cmp(big.NewInt(0)) == 0 {
				g = g.Mul(g, n)
				a = a.Div(a, n)
				b = b.Div(b, n)

			} else {
				break
			}
		}
		n = n.Add(n, big.NewInt(1))
	}

	return g
}
