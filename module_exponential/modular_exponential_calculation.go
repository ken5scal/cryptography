package module_exponential

import (
	"errors"
	"math"
	"math/big"
)

/*
RSA Encryption/Decryption is done by Modular Exponential Calculation(MEC) <- 冪剰余計算
S = a^m mod N
 Encryption:
  - S: Cipher Text(C)
  - a: Plain Text(M)
  - m: Public Exponent公開指数(e)
  - N: Public Modulus(N)公開モジュラス
  - (m, N) = (e, N): public key
  - C = M ^ e mod N
 Decryption:
  - S: Plain Text(M)
  - a: Cipher Text(C)
  - m: Secret Exponent秘密指数(d)
  - N: Public Modulus(N)公開モジュラス
  - (m, N) = (d, N): private key
  - M = C ^ d mod N
*/

// ChineseRemainderTheorem uses prime factors(p, q) to calculate MEC
// its pre-requisite is to know p and q. So it can only utilized in DECRYPTION.
func ChineseRemainderTheorem(c, p, q, dp, dq, v *big.Int) (M *big.Int, err error) {
	// Initialize
	cp := new(big.Int)
	cq := new(big.Int)
	V := new(big.Int)
	M = new(big.Int)

	cp.Set(c)
	cp.Mod(c, p)
	cq.Set(c)
	cq.Mod(c, q)

	Mp, err := ModPow(c, dp, p)
	if err != nil {
		return nil, err
	}

	Mq, err := ModPow(c, dq, q)
	if err != nil {
		return nil, err
	}

	// V = v(Mq - Mp) mod q
	V.Mul(v, V.Sub(Mq, Mp)).Mod(V, q)
	for V.Int64() < 0 {
		V.Add(V, q)
	}

	// M = Vp + Mp
	M.Add(M.Mul(V, p), Mp)
	return M, err
}

// ModPowSlidingWindow is another method to calculate MEC using Sliding Window Mod Pow
func ModPowSlidingWindow(a, m, N *big.Int, w int) (s *big.Int, err error) {
	if a.Sign() <= 0 || N.Sign() <= 0 || m.Sign() <= 0 || w <= 0 {
		return nil, errors.New("Input must be positive number")
	}

	at := makeDataTableForSlidingWindow(a.Int64(), N.Int64(), int64(w))
	s = big.NewInt(1)
	for j := m.BitLen() - 1; j >= 0; {
		if m.Bit(j) == 0 {
			s.Exp(s, big.NewInt(2), N)
			j--
		} else {
			l := int(math.Max(float64(j-w+1), 0))
			for ; j > l; l++ {
				if m.Bit(l) != 0 {
					break
				}
			}

			mjl := int64(0)
			for i := j; i >= l; i-- {
				mjl <<= 1
				if m.Bit(i) != 0 {
					mjl |= 1
				}
				s.Exp(s, big.NewInt(2), N)
			}
			s.Mul(s, at[mjl>>1]).Mod(s, N)
			j = l - 1
		}
	}
	return s, nil
}

// makeDataTableForSlidingWindow makes sliding window table for ModPowSlidingWindow
func makeDataTableForSlidingWindow(a, N, w int64) (at []*big.Int) {
	length := 1 << uint(w-1)
	at = make([]*big.Int, length)
	b := a * a % N
	at[0] = big.NewInt(a % N)
	for j := 1; j < length; j++ {
		at[j] = big.NewInt(at[j-1].Int64() * b % N)
	}
	return at
}

// ModPow2wary is another method to do MEC using Window(2w-ary) ModPow
func ModPow2wary(a, m, N *big.Int, w int) (*big.Int, error) {
	if a.Sign() <= 0 || N.Sign() <= 0 || m.Sign() <= 0 || w <= 0 {
		return nil, errors.New("Input must be positive number")
	}

	length := 1 << uint(w)
	table := make([]*big.Int, length)
	table[0] = big.NewInt(1)
	for k := 1; k < length; k++ {
		table[k] = big.NewInt(table[k-1].Int64() * a.Int64() % N.Int64())
	}

	S := big.NewInt(1)
	for j := (m.BitLen()+w-1)/w - 1; j >= 0; j-- {
		for i := 0; i < w; i++ {
			S.Mul(S, S).Mod(S, N)
		}

		mjw := int64(0)
		for i := w - 1; i >= 0; i-- {
			mjw <<= 1
			if m.Bit(j*w+i) != 0 {
				mjw |= 1
			}
		}
		S.Mul(S, table[mjw]).Mod(S, N)
	}

	return S, nil
}

// ModPow calculates the most fundamental binary function.
func ModPow(a, m, N *big.Int) (*big.Int, error) {
	if N.Sign() <= 0 || m.Sign() <= 0 {
		return nil, errors.New("Input must be positive number")
	}

	S := big.NewInt(1)

	for j := m.BitLen() - 1; j >= 0; j-- {
		S.Mul(S, S).Mod(S, N)

		if m.Bit(j) == 1 {
			S.Mul(S, a).Mod(S, N)
		}
	}
	return S, nil
}
