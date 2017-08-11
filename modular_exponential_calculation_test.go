package main

import (
	"fmt"
	"math/big"
	"testing"
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

}