package main

import (
	"fmt"
	"math/big"
	"testing"
)

var a = big.NewInt(1024)
var b = big.NewInt(15000)
var expected = big.NewInt(8)

var testSets = []struct {
	a, b, expectedX, expectedY *big.Int
} {
	{big.NewInt(79), big.NewInt(176), big.NewInt(-49), big.NewInt(22)},
	{big.NewInt(2793), big.NewInt(828), big.NewInt(67), big.NewInt(-226)},
}

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
	for _, tt := range testSets {
		x, y, r, err := ExtendedPublicGCD(tt.a, tt.b)
		if err != nil {
			t.Errorf("Error: %v\n", err)
			return
		}
		expectedR, _ := BinaryEuclidGCD(tt.a, tt.b)

		if tt.expectedX.Cmp(x) != 0 {
			t.Error(fmt.Sprintf("expected x as %v, but got %v", tt.expectedX, x))
		} else if tt.expectedY.Cmp(y) != 0 {
			t.Error(fmt.Sprintf("expected y as %v, but got %v", tt.expectedY, y))
		} else if expectedR.Cmp(r) != 0 {
			t.Error(fmt.Sprintf("expected r as %v, but got %v", expectedR, r))
		}
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
