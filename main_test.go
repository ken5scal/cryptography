package main

import (
	"fmt"
	"math/big"
	"testing"
)

var a = big.NewInt(1024)
var b = big.NewInt(15000)
var expected = big.NewInt(8)

func TestGcd(t *testing.T) {
	actual, _ := Gcd(a, b)
	if expected.Cmp(actual) != 0 {
		t.Error(fmt.Sprintf("expected %v, but got %v", expected, actual))
	}
}

func TestBinaryEuclidGCD(t *testing.T) {
	actual, _ := BinaryEuclidGCD(a, b)
	if expected.Cmp(actual) != 0 {
		t.Error(fmt.Sprintf("expected %v, but got %v", expected, actual))
	}
}

func TestBadBinaryEuclidGCD(t *testing.T) {
	badA := big.NewInt(-1)
	actual, err := BinaryEuclidGCD(badA, b)
	if actual != nil && err == nil {
		t.Error("actual must be null value. Error must not be a null value")
	}
}

func TestExtendedPublicGCD(t *testing.T) {
	a := big.NewInt(2793)
	b := big.NewInt(828)
	expectedX := big.NewInt(67)
	expectedY := big.NewInt(-226)
	expectedR := big.NewInt(3)
	
	x, y, r, err := ExtendedPublicGCD(a, b)
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}

	if expectedX.Cmp(x) != 0 {
		t.Error(fmt.Sprintf("expected x as %v, but got %v", expectedX, x))
	} else if expectedY.Cmp(y) != 0 {
		t.Error(fmt.Sprintf("expected y as %v, but got %v", expectedY, y))
	} else if expectedR.Cmp(r) != 0 {
		t.Error(fmt.Sprintf("expected r as %v, but got %v", expectedR, r))
	}
}

func BenchmarkGcd_normal(bench *testing.B) {
	for i := 0; i < 5000; i++ {
		Gcd(a, b)
	}
}

func BenchmarkGCD_binaryEuclid(bench *testing.B) {
	for i := 0; i < 5000; i++ {
		BinaryEuclidGCD(a, b)
	}
}
