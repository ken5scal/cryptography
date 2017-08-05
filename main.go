package main

import (
	"errors"
	"fmt"
	"math/big"

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

	a = big.NewInt(1024)
	b = big.NewInt(15000)
	g, _ = BinaryEuclidGCD(a, b)
	fmt.Println(g)

	a = big.NewInt(2793)
	b = big.NewInt(828)
	x, y, g, _ := ExtendedPublicGCD(a, b)
	fmt.Println(x, y, g)
}

// BinaryEuclidGCD calculated gcd(最大公約数) using Binaryユークリッド互除法アルゴリズム
func BinaryEuclidGCD(aOrig, bOrig *big.Int) (*big.Int, error) {
	var zero big.Int
	if aOrig.Cmp(&zero) != 1 || bOrig.Cmp(&zero) != 1 {
		return nil, errors.New("Arguments must be positive number.")
	}

	a, b := CopyData(aOrig, bOrig)
	g := big.NewInt(1)

	isEven := func(num int64) bool { return num%2 == 0 }
	calcT := func(a, b int64) *big.Int {
		t := (a - b) / 2
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
		} else if isEven(a.Int64()) && !isEven(b.Int64()) {
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

	return g.Mul(g, b), nil
}

// EuclidGCD calculated gcd(最大公約数) using ユークリッド互除法アルゴリズム
func EuclidGCD(aOrig, bOrig *big.Int) *big.Int {
	a, b := CopyData(aOrig, bOrig)

	for r := big.NewInt(0); ; {
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
func Gcd(aOrig, bOrig *big.Int) (*big.Int, error) {
	a, b := CopyData(aOrig, bOrig)
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
		for {
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

func ExtendedPublicGCD(aOrig, bOrig *big.Int) (x, y, r *big.Int, err error) {
	a, b := CopyData(aOrig, bOrig)
	if a.Sign() <= 0 || b.Sign() <= 0 {
		return nil, nil, nil, errors.New("Input must be positive number")
	}

	// Initializing
	xPrev := big.NewInt(1)
	yPrev := big.NewInt(0)
	rPrev := big.NewInt(a.Int64())

	x = big.NewInt(0)
	y = big.NewInt(1)
	r = big.NewInt(b.Int64())

	xNext := big.NewInt(0)
	yNext := big.NewInt(0)
	rNext := big.NewInt(0)
	qNext := big.NewInt(0)

	for j := 0; r.Int64() != 0; j++ {
		qNext.Div(rPrev, r)
		rNext.Mod(rPrev, r)

		xNext.Sub(xPrev, xNext.Mul(qNext, x))
		yNext.Sub(yPrev, yNext.Mul(qNext, y))

		xPrev.Set(x); yPrev.Set(y); rPrev.Set(r)
		x.Set(xNext); y.Set(yNext); r.Set(rNext)
	}

	return xPrev, yPrev, rPrev, nil
}

func CopyData(aOrig, bOrig *big.Int) (a, b *big.Int) {
	// Making sure immutability
	a = big.NewInt(aOrig.Int64())
	b = big.NewInt(bOrig.Int64())
	return a, b
}
