package main

import (
	"fmt"
	"math/big"
	"testing"
)

func TestIsPrimeByMillerRabinTest(t *testing.T) {
	r := big.NewInt(7)
	t2 := big.NewInt(1)
	actual, err := IsPrimeByMillerRabinTest(r, t2)
	if err != nil {
		t.Error(err)
	}

	if !actual {
		t.Errorf("Expected r %v to be prime, but not.\n", r)
	}
}

func TestFindSandK(t *testing.T) {
	expectedS := big.NewInt(3)
	expectedK := big.NewInt(11)
	s, k := findSandK(big.NewInt(89))
	eMessage := ""
	if expectedS.Cmp(s) != 0 {
		eMessage = fmt.Sprintf("Expected s %v, but was %v\n", expectedS, s)
	}
	if expectedK.Cmp(k) != 0 {
		eMessage = fmt.Sprintf("Expected k %v, but was %v\n", expectedK, k)
	}

	if eMessage != "" {
		t.Errorf(eMessage)
	}
}

func TestIsPrimeByFermat(t *testing.T) {
	r := big.NewInt(7)
	t2 := big.NewInt(1)
	actual, err := IsPrimeByFermatTest(r, t2)
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

	_, err := IsPrimeByFermatTest(zero, three)
	if err == nil {
		t.Errorf("Expected to return error due to illegal argument.")
	}

	_, err = IsPrimeByFermatTest(three, zero)
	if err == nil {
		t.Errorf("Expected to return error due to illegal argument.")
	}
}
