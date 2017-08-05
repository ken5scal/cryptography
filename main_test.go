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
	actual, _ := BinaryEuclidGCD(a, b)
	if expected.Cmp(actual) != 0 {
		t.Error(fmt.Sprintf("expected %v, but got %v", expected, actual))
	}
}

func TestBadBinaryEuclidGCD(t *testing.T) {
	a = big.NewInt(-1)
	actual, err := BinaryEuclidGCD(a, b)
	if actual != nil && err == nil {
		t.Error("actual must be null value. Error must not be a null value")
	}
}

func BenchmarkBinaryEuclidGCD(bench *testing.B) {
	bench.ResetTimer()
	BinaryEuclidGCD(a,b)
}