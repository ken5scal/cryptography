package main

import (
	"testing"
	"math/big"
	"fmt"
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
	var a = big.NewInt(1024)
	var b = big.NewInt(15000)
	var expected = big.NewInt(8)
	actual, _ := BinaryEuclidGCD(a, b)
	if expected.Cmp(actual) != 0 {
		t.Error(fmt.Sprintf("expected %v, but got %v", expected, actual))
	}
}

func TestBadBinaryEuclidGCD(t *testing.T) {
	var b = big.NewInt(15000)
	badA := big.NewInt(-1)
	actual, err := BinaryEuclidGCD(badA, b)
	if actual != nil && err == nil {
		t.Error("actual must be null value. Error must not be a null value")
	}
}

func BenchmarkGcd_normal(bench *testing.B) {
	for i := 0; i < 10000; i++ {
		Gcd(a, b)
	}
}

func BenchmarkGCD_binaryEuclid(bench *testing.B) {
	for i := 0; i < 10000; i++ {
		BinaryEuclidGCD(a, b)
	}
}