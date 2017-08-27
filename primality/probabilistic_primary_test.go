package primality

import (
	"fmt"
	"math/big"
	"testing"
)

var primes = []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
var composites = []int64{9, 15, 21, 25, 27, 33, 35, 39, 45, 49, 51, 55, 57, 63, 65, 69, 75, 77, 81, 85, 87, 91, 93, 95, 99}

func TestIsPrimeByMillerRabinTest(t *testing.T) {
	t2 := big.NewInt(30)

	for _, p := range primes {
		isPrime, err := IsPrimeByMillerRabinTest(big.NewInt(p), t2)
		if err != nil {
			t.Error(err)
		}

		if !isPrime {
			t.Errorf("Expected %v to be prime, but not.\n", p)
		}
	}
}

func TestIsPrimeByMillerRabinTest2(t *testing.T) {
	t2 := big.NewInt(30)

	for _, composite := range composites {
		isPrime, err := IsPrimeByMillerRabinTest(big.NewInt(composite), t2)
		if err != nil {
			t.Error(err)
		}

		if isPrime {
			t.Errorf("Expected %v to be prime, but not.\n", composite)
		}
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
