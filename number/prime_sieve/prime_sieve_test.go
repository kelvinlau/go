package prime_sieve

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
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

func testEquals(a, b []int, t *testing.T) {
	if len(a) != len(b) {
		t.Fatalf("Wrong length: %d", len(a))
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Fatalf("Wrong element %d: got %d, want %d", i, a[i], b[i])
		}
	}
}

func TestFactorize(t *testing.T) {
	test := func(n int) {
		fs := Factorize(n)
		m := 1
		for _, f := range fs {
			if !IsPrime(f.p) {
				t.Fatalf("Factorize %d, got factor %d which is not a prime.", n, f.p)
			}
			for k := 0; k < f.k; k++ {
				m *= f.p
			}
		}
		if n != m {
			t.Fatalf("Product of all factor = %d, expected %d.", m, n)
		}
	}

	for _, n := range []int{1, 2, 3, 7, 9, 10, 14, 16, 100, 23, 49, 101, 2000, 100000} {
		test(n)
	}
}
