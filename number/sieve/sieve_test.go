package sieve

import (
	"testing"
)

func TestSmallPrimes(t *testing.T) {
	es := []int{2, 3, 5, 7, 11, 13, 17, 19}
	ps := []int{}
	for i := 1; i < 20; i++ {
		if IsPrime(i) {
			ps = append(ps, i)
		}
	}
	t.Logf("ps: %#v", ps)
	testEquals(ps, es, t)
}

func TestBigPrimes(t *testing.T) {
	ps := []int{1<<30 - 35, 1<<30 - 41, 1<<25 - 141}
	qs := []int{1<<30 - 37, 1<<30 - 43, 1<<25 - 163}
	for _, p := range ps {
		if !IsPrime(p) {
			t.Errorf("%d should be a prime.", p)
		}
	}
	for _, q := range qs {
		if IsPrime(q) {
			t.Errorf("%d should not be a prime.", q)
		}
	}
}

func testEquals(a, b []int, t *testing.T) {
	if len(a) != len(b) {
		t.Fatalf("Wrong length: %d", len(a))
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Errorf("Wrong element %d: got %d, want %d", i, a[i], b[i])
		}
	}
}

func TestFactorize(t *testing.T) {
	test := func(n int) {
		fs := Factorize(n)
		m := 1
		for _, f := range fs {
			if !IsPrime(f.P) {
				t.Errorf("Factorize %d, got factor %d which is not a prime.", n, f.P)
			}
			for k := 0; k < f.K; k++ {
				m *= f.P
			}
		}
		if n != m {
			t.Errorf("Product of all factor = %d, expected %d.", m, n)
		}
	}

	for _, n := range []int{1, 2, 3, 7, 9, 10, 14, 16, 100, 23, 49, 101, 2000, 100000, 481239000, 1<<30 - 35} {
		test(n)
	}
}
