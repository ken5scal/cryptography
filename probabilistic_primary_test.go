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
	zero := big.NewInt(0)
	three := big.NewInt(10)

	_, err := IsPrimeByFermat(zero, three)
	if err == nil {
		t.Errorf("Expected to return error due to illegal argument.")
	}

	_, err = IsPrimeByFermat(three, zero)
	if err == nil {
		t.Errorf("Expected to return error due to illegal argument.")
	}
}
