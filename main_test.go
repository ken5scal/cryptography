package main

import (
	"testing"
	"math/big"
	"fmt"
)

func TestGcd(t *testing.T) {
	a := big.NewInt(780)
	b := big.NewInt(600)
	expected := big.NewInt(60)
	actual, _ := Gcd(a, b)
	if expected.Cmp(actual) != 0 {
		t.Error(fmt.Sprintf("expected %v, but got %v", expected, actual))
	}
}
