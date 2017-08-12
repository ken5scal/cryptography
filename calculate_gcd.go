package main

import (
	"math/big"
	"errors"
)

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

// ExtendedEuclidGCD
func ExtendedEuclidGCD(aOrig, bOrig *big.Int) (x, y, r *big.Int, err error) {
	a, b := CopyData(aOrig, bOrig)
	if a.Sign() <= 0 || b.Sign() <= 0 {
		return nil, nil, nil, errors.New("Input must be positive number")
	}

	// Initializing
	xPrev := big.NewInt(1)
	x = big.NewInt(0)
	xNext := big.NewInt(0)
	yPrev := big.NewInt(0)
	y = big.NewInt(1)
	yNext := big.NewInt(0)
	rPrev := big.NewInt(a.Int64())
	r = big.NewInt(b.Int64())
	rNext := big.NewInt(0)
	qNext := big.NewInt(0)

	for r.Int64() != 0 {
		qNext.Div(rPrev, r)
		rNext.Mod(rPrev, r)

		xNext.Sub(xPrev, xNext.Mul(qNext, x))
		yNext.Sub(yPrev, yNext.Mul(qNext, y))

		xPrev.Set(x)
		x.Set(xNext)

		yPrev.Set(y)
		y.Set(yNext)

		rPrev.Set(r)
		r.Set(rNext)
	}

	return xPrev, yPrev, rPrev, nil
}

// ModInverse is basically ExtendedEuclidGCD adapted to RSA
func ModInverse(e, l *big.Int) (*big.Int, error) {
	if e.Sign() <= 0 || l.Sign() <= 0 {
		return nil, errors.New("Input must be positive number")
	}

	dPrev := big.NewInt(1)
	d := big.NewInt(0)
	dNext := big.NewInt(0)

	rPrev := big.NewInt(e.Int64())
	r := big.NewInt(l.Int64())
	rNext := big.NewInt(0)

	xNext := big.NewInt(0)

	for r.Int64() != 0 {
		dNext.Div(rPrev, r)
		rNext.Mod(rPrev, r)

		xNext.Sub(dPrev, xNext.Mul(dNext, d))

		dPrev.Set(d)
		d.Set(xNext)

		rPrev.Set(r)
		r.Set(rNext)
	}

	return dPrev.Mod(dPrev, l), nil
}

func CopyData(aOrig, bOrig *big.Int) (a, b *big.Int) {
	// Making sure immutability
	a = big.NewInt(aOrig.Int64())
	b = big.NewInt(bOrig.Int64())
	return a, b
}
