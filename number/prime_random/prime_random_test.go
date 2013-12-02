package prime_random

import (
	"math/rand"
	"testing"

	ps "github.com/kelvinlau/go/number/prime_sieve"
)

func TestPollarRho(t *testing.T) {
	test := func(n int64) {
		d := PollarRho(n)
		if d <= 1 || d >= n || n%d != 0 {
			t.Fatalf("PollarRho(%d) got %d.", n, d)
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
			t.Fatalf("MillerRabin(%d) got %v, expected %v.", n, g, e)
		}
	}

	for i := 0; i < 100; i++ {
		n := int(rand.Int31())
		test(int64(n), ps.IsPrime(n))
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
				t.Fatalf("Non-prime factor %d of %d.", p, n)
			}
			m *= p
		}
		if n != m {
			t.Fatalf("Product of %v = %d, != %d.", ps, m, n)
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
}
