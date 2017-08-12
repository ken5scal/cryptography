package main

import (
	"fmt"
	"math/big"
	"testing"
	crand "crypto/rand"
	"encoding/binary"
)

func TestChineseRemainderTheorem(t *testing.T) {
	expectedS := big.NewInt(0)
	C := big.NewInt(52)
	p := big.NewInt(7)
	q := big.NewInt(11)
	d := big.NewInt(13)

	dp := big.NewInt(0)
	dq := big.NewInt(0)
	v := big.NewInt(0)

	dp.Mod(d, dp.Sub(p, big.NewInt(1)))
	dq.Mod(d, dq.Sub(q, big.NewInt(1)))
	v.ModInverse(p, q)

	s, err := ChineseRemainderTheorem(C, p, q, dp, dq, v)

	N := big.NewInt(p.Int64() * q.Int64())
	expectedS.Exp(C, d, N)
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}

	if expectedS.Cmp(s) != 0 {
		t.Error(fmt.Sprintf("expected %v, but got %v", expectedS, s))
	}
}

func TestMakeDataTableForSlidingWindow(t *testing.T) {
	w := 4
	length := 1 << uint(w-1)
	at := makeDataTableForSlidingWindow(10, 19, int64(w))
	expected := []*big.Int{big.NewInt(10), big.NewInt(12), big.NewInt(3),
		big.NewInt(15), big.NewInt(18), big.NewInt(14),big.NewInt(13), big.NewInt(8)}
	if len(at) != length {
		t.Errorf("Length should be: %v, but was %v\n", length, len(at))
		return
	}
	for i, v := range expected {
		if at[i].Cmp(v) != 0 {
			t.Errorf("Table at index %v should be %v, but was %v\n", i, v, at[i])
		}
	}
}

func TestModPowSlidingWindow(t *testing.T) {
	expectedS := big.NewInt(0)

	a := big.NewInt(10)
	m := big.NewInt(2405)
	n := big.NewInt(19)
	expectedS.Exp(a, m, n)

	s, err := ModPowSlidingWindow(a, m, n, 4)
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}
	if expectedS.Cmp(s) != 0 {
		t.Error(fmt.Sprintf("expected %v, but got %v", expectedS, s))
	}
}

func TestModPow2wary(t *testing.T) {
	expectedS := big.NewInt(0)

	a := big.NewInt(10)
	m := big.NewInt(2405)
	n := big.NewInt(19)
	expectedS.Exp(a, m, n)
	s, err := ModPow2wary(a, m, n, 4)
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}
	if expectedS.Cmp(s) != 0 {
		t.Error(fmt.Sprintf("expected %v, but got %v", expectedS, s))
	}
}

func TestModPow(t *testing.T) {
	expectedS := big.NewInt(0)

	a := big.NewInt(10)
	m := big.NewInt(2405)
	n := big.NewInt(19)
	expectedS.Exp(a, m, n)

	s, err := ModPow(a, m, n)
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}
	if expectedS.Cmp(s) != 0 {
		t.Error(fmt.Sprintf("expected %v, but got %v", expectedS, s))
	}
}

// ToDo Requires Improvement: Make 2048 fixed length of random number
func BenchmarkModPowSlidingWindow(b *testing.B) {

	bitCount := 2048
	dataCount := 10
	windowSize := 6

	randoms := make([]*big.Int, dataCount + 2)
	k := make([]byte, bitCount)

	for i := 0; i < len(randoms); {
		crand.Read(k)
		by, _ := binary.Varint(k)

		if by > 0 {
			randoms[i] = big.NewInt(by)
			i++
		}
	}

	for i := 0; i < dataCount; i++ {
		ModPowSlidingWindow(randoms[i], randoms[i+1], randoms[i+2], windowSize)
	}
}

// ToDo Requires Improvement: Make 2048 fixed length of random number
func BenchmarkModPow(b *testing.B) {

	bitCount := 2048
	dataCount := 10

	randoms := make([]*big.Int, dataCount + 2)
	k := make([]byte, bitCount)

	for i := 0; i < len(randoms); {
		crand.Read(k)
		by, _ := binary.Varint(k)

		if by > 0 {
			randoms[i] = big.NewInt(by)
			i++
		}
	}

	for i := 0; i < dataCount; i++ {
		ModPow(randoms[i], randoms[i+1], randoms[i+2])
	}
}