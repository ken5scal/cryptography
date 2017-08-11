package main

import (
	"fmt"
	"math/big"
	"testing"
	crand "crypto/rand"
	"encoding/binary"
)

func TestMakeDataTableForSlidingWindow(t *testing.T) {
	at := makeDataTableForSlidingWindow(10, 19, 4)
	fmt.Println(at)
}

func TestModPowSlidingWindow(t *testing.T) {
	expectedS := big.NewInt(0)

	a := big.NewInt(10)
	m := big.NewInt(2405)
	n := big.NewInt(19)
	expectedS.Exp(a, m, n)

	fmt.Printf("a: %v, m: %v, n: %v\n", a, m, n)
	s, err := ModPowSlidingWindow(a, m, n, 4)
	if err != nil {
		t.Errorf("Error: %v\n", err)
		return
	}
	fmt.Println(s, expectedS)
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