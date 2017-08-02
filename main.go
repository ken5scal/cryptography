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

	a = big.NewInt(135632)
	b = big.NewInt(44461)
	g = EuclidGCD(a, b)
	fmt.Println(g)

	a = big.NewInt(646)
	b = big.NewInt(408)
	g = BinaryEuclidGCD(a, b)
	fmt.Println(g)
}

// BinaryEuclidGCD calculated gcd(最大公約数) using Binaryユークリッド互除法アルゴリズム
func BinaryEuclidGCD(aOrig, bOrig *big.Int) *big.Int {
	// Making sure immutability
	a := big.NewInt(aOrig.Int64())
	b := big.NewInt(bOrig.Int64())
	g := big.NewInt(1)

	isEven := func(num int64) bool {return num % 2 == 0}
	calcT := func(a, b int64) *big.Int {
		t := (a-b)/2
		if t < 0 {
			t *= -1
		}
		return big.NewInt(t)
	}
	two := big.NewInt(2)

	for a.Int64() > 0 {
		if isEven(a.Int64()) && isEven(b.Int64()) {
			a.Div(a, two)
			b.Div(b, two)
			g.Mul(g, two)
		} else if isEven(a.Int64()) && !isEven(b.Int64()){
			a.Div(a, two)
		} else if !isEven(a.Int64()) && isEven(b.Int64()) {
			b.Div(b, two)
		} else {
			t := calcT(a.Int64(), b.Int64())
			if a.Cmp(b) >= 0 {
				a = t
				continue
			}
			b = t
		}
	}

	return g.Mul(g, b)
}

// EuclidGCD calculated gcd(最大公約数) using ユークリッド互除法アルゴリズム
func EuclidGCD(aOrig, bOrig *big.Int) *big.Int {
	// Making sure immutability
	a := big.NewInt(aOrig.Int64())
	b := big.NewInt(bOrig.Int64())

	for r := big.NewInt(0);; {
		r.Mod(a, b)
		if r.Int64() == 0 {
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
