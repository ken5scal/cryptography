package main

import (
	"fmt"
	"math/big"
	"testing"
)

func TestGetWindow(t *testing.T) {
	m := big.NewInt(1501)
	GetWindow(m, 3)
}

func TestModPow2wary(t *testing.T) {
	expectedS := big.NewInt(0)

	a := big.NewInt(10)
	m := big.NewInt(13)
	n := big.NewInt(19)
	expectedS.Exp(a, m, n)
	s, err := ModPow2wary(a, m, n, 3)
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
	m := big.NewInt(13)
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
