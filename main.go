package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(780)
	b := big.NewInt(600)
	g, _ := Gcd(a, b)
	fmt.Println(g)

	a = big.NewInt(135632)
	b = big.NewInt(44461)
	g = EuclidGCD(a, b)
	fmt.Println(g)

	a = big.NewInt(1024)
	b = big.NewInt(15000)
	g, _ = BinaryEuclidGCD(a, b)
	fmt.Println(g)

	a = big.NewInt(2793)
	b = big.NewInt(828)
	x, y, g, _ := extendedEuclidGCD(a, b)
	fmt.Println(x, y, g)
}