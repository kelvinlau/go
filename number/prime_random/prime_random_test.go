package prime_random

import (
	"math/rand"
	"testing"

	"github.com/kelvinlau/go/number/sieve"
)

func TestPollarRho(t *testing.T) {
	test := func(n int64) {
		d := PollarRho(n)
		if d <= 1 || d >= n || n%d != 0 {
			t.Errorf("PollarRho(%d) got %d.", n, d)
		} else {
			t.Logf("PollarRho(%d) got %d.", n, d)
		}
	}

	test(20)
	test(128)
	test(28911)
	test(1<<48 - 1)
	test(1<<50 - 1)
	test(1<<52 - 1)
	test(1<<54 - 1)
}

func TestMillerRabin(t *testing.T) {
	test := func(n int64, e bool) {
		if g := MillerRabin(n); g != e {
			t.Errorf("MillerRabin(%d) got %v, expected %v.", n, g, e)
		}
	}

	for i := 0; i < 100; i++ {
		n := int(rand.Int31())
		test(int64(n), sieve.IsPrime(n))
	}
	test(123456789123456, false)
	test(1<<48-257, true)
	test(1<<60-171, false)
	test(1<<60-173, true)
	test(1<<61-1, true)
}

func TestFactorize(t *testing.T) {
	test := func(n int64) {
		ps := Factorize(n)
		t.Logf("Factorize(%d) = %v.", n, ps)
		m := int64(1)
		for _, p := range ps {
			if !MillerRabin(p) {
				t.Errorf("Non-prime factor %d of %d.", p, n)
			}
			m *= p
		}
		if n != m {
			t.Errorf("Product of %v = %d, != %d.", ps, m, n)
		}
	}

	for n := int64(1); n < 1000; n++ {
		test(n)
	}
	test(123456789123456)
	test(1<<48 - 257)
	test(1<<60 - 171)
	test(1<<60 - 173)
	test(1<<61 - 1)
	test((1<<28 - 57) * (1<<28 - 89))
}

func TestDivisors(t *testing.T) {
	test := func(n int64) {
		ds := Divisors(n)
		t.Logf("Divisors(%d) = %v.", n, ds)
		for _, d := range ds {
			if n%d != 0 {
				t.Errorf("%d is not dividible by %d.", n, d)
			}
		}
	}

	for n := int64(1); n < 1000; n++ {
		test(n)
	}
	test(123456789123456)
	test(1<<48 - 257)
	test(1<<60 - 171)
	test(1<<60 - 173)
	test(1<<61 - 1)
	test((1<<28 - 57) * (1<<28 - 89))
}

func BenchmarkFactorize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorize(rand.Int63n(1 << 61))
	}
}
