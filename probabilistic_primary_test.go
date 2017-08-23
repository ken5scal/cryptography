package main

import (
	"math/big"
	"testing"
)

func TestIsPrimeByFermat(t *testing.T) {
	r := big.NewInt(7)
	t2 := big.NewInt(1)
	actual, err := IsPrimeByFermat(r, t2)
	if err != nil {
		t.Error(err)
	}

	if !actual {
		t.Errorf("Expected r %v to be prime, but not.\n", r)
	}
}

func TestIllegalArgumentPrimeByFermat(t *testing.T) {
	r := big.NewInt(0)
	t2 := big.NewInt(1)

	_, err := IsPrimeByFermat(r, t2)

	if err != nil {
		t.Errorf("Expected to return error due to illegal argument.")
	}
}
