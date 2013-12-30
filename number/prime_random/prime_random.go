// Package prime_random implements prime number checker using random algorithms.
package prime_random

import (
	"math/rand"
	"sort"

	"github.com/kelvinlau/go/number"
)

// PollarRho returns a divisor of n.
func PollarRho(n int64) int64 {
	if n%2 == 0 {
		return 2
	}
	for {
		x := rand.Int63n(n)
		y := x
		c := rand.Int63n(n)
		for {
			f := func(x int64) int64 {
				return (number.ModularMultiply(x, x, n) + c) % n
			}
			x = f(x)
			y = f(f(y))
			d := number.GCD(abs(x-y), n)
			if d == n {
				break
			}
			if d > 1 {
				return d
			}
		}
	}
}

// MillerRabin test if n is a prime.
func MillerRabin(n int64) bool {
	return MillerRabinK(n, 40)
}

// MillerRabinK test if n is a prime, testing for k times.
func MillerRabinK(n int64, k int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%2 == 0 {
		return false
	}
	d := n - 1
	s := 0
	for d%2 == 0 {
		d /= 2
		s++
	}

	for ; k > 0; k-- {
		x := number.ModularPower(rand.Int63n(n-3)+2, d, n)
		if x == 1 || x == n-1 {
			continue
		}

		r := 1
		for ; r < s; r++ {
			x = number.ModularMultiply(x, x, n)
			if x == 1 {
				return false
			}
			if x == n-1 {
				break
			}
		}
		if r == s {
			return false
		}
	}
	return true
}

// Factorize factorizes n, returns all prime factors.
func Factorize(n int64) (fs []int64) {
	if n == 1 {
		return
	}
	if MillerRabin(n) {
		fs = append(fs, n)
	} else {
		d := PollarRho(n)
		fs = append(fs, Factorize(d)...)
		fs = append(fs, Factorize(n/d)...)
	}
	return
}

// Divisors return all divisors of n.
func Divisors(n int64) (ds []int64) {
	fs := Factorize(n)
	sort.Sort(factors(fs))

	ds = append(ds, 1)
	for i, j := 0, 0; i < len(fs); i = j {
		for j < len(fs) && fs[i] == fs[j] {
			j++
		}
		p := fs[i]
		k := j - i
		o := len(ds)
		for z := 0; z < k*o; z++ {
			ds = append(ds, ds[len(ds)-o]*p)
		}
	}
	return
}

type factors []int64

func (fs factors) Len() int           { return len(fs) }
func (fs factors) Swap(i, j int)      { fs[i], fs[j] = fs[j], fs[i] }
func (fs factors) Less(i, j int) bool { return fs[i] < fs[j] }

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}
