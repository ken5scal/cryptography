package cryptography

import "math/big"

func Gcd(a, b *big.Int) big.Int {
	g := big.NewInt(1)
	if a.Cmp(b) < 0 {
		tmp := a; a = b; b= tmp
	}

	n := big.NewInt(2)
	for n.Cmp(b) < 0 {
		for a.Mod(a, n) == big.NewInt(0) && b.Mod(b, n) == big.NewInt(0){
			g = g.Mul(g, n)
			a = a.Div(a, n)
			b = b.Div(b, n)
		}
		n = n.Add(n, big.NewInt(1))
	}

	return g
}
