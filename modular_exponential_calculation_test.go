package main

import (
	"testing"
	"math/big"
	"fmt"
)

func TestModPow(t *testing.T) {
	a := big.NewInt(2048)
	m := big.NewInt(10)
	n := big.NewInt(6)

	expectedS := big.NewInt(7148)
	s := ModPow(*a,*m,*n)
	if expectedS.Cmp(&s) != 0 {
		t.Error(fmt.Sprintf("expected %v, but got %v", expected, s))
	}
}
