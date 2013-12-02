package ptree

import (
	"math/rand"
	"testing"
)

func TestSmall(t *testing.T) {
	a := []int{3, 6, 6, 2, 8, 4, 2}
	n := len(a)
	pt := New(a)
	testRLB(t, pt, 0, n, 1, 2)
	testRLB(t, pt, 0, n, 2, 2)
	testRLB(t, pt, 0, n, 3, 3)
	testRLB(t, pt, 0, n, 4, 4)
	testRLB(t, pt, 0, n, 5, 6)
	testRLB(t, pt, 0, n, 6, 6)
	testRLB(t, pt, 0, n, 7, 8)
	testRLB(t, pt, 0, n, 8, 8)
	testRLB(t, pt, 0, n, 9, Inf)
	testRLB(t, pt, 0, 4, 1, 2)
	testRLB(t, pt, 0, 4, 2, 2)
	testRLB(t, pt, 0, 4, 3, 3)
	testRLB(t, pt, 0, 4, 4, 6)
	testRLB(t, pt, 0, 4, 5, 6)
	testRLB(t, pt, 0, 4, 6, 6)
	testRLB(t, pt, 0, 4, 7, Inf)
	testRLB(t, pt, 1, 6, 1, 2)
	testRLB(t, pt, 1, 6, 2, 2)
	testRLB(t, pt, 1, 6, 3, 4)
	testRLB(t, pt, 1, 6, 4, 4)
	testRLB(t, pt, 1, 6, 5, 6)
	testRLB(t, pt, 1, 6, 6, 6)
	testRLB(t, pt, 1, 6, 7, 8)
	testRLB(t, pt, 1, 6, 8, 8)
	testRLB(t, pt, 1, 6, 9, Inf)
}

func TestLarge(t *testing.T) {
	n := 1000
	a := make([]int, n)
	lim := 10000
	for i := range a {
		a[i] = rand.Intn(lim)
	}
	pt := New(a)
	for k := 0; k < 1000; k++ {
		l := rand.Intn(n)
		r := rand.Intn(n)
		if l < r {
			l, r = r, l
		}
		x := rand.Intn(lim)
		y := Inf
		for i := l; i < r; i++ {
			if a[i] >= x && a[i] < y {
				y = a[i]
			}
		}
		testRLB(t, pt, l, r, x, y)
	}
}

func testRLB(t *testing.T, pt *PTree, l, r, v, e int) {
	if g := pt.RangeLB(l, r, v); g != e {
		t.Fatalf("RLB of %d in [%d, %d) should be %d, got %d.", v, l, r, e, g)
	}
}

func BenchmarkPTree(b *testing.B) {
	n := 100000
	a := make([]int, n)
	lim := 1000000
	for i := range a {
		a[i] = rand.Intn(lim)
	}
	pt := New(a)
	b.ResetTimer()
	for k := 0; k < b.N; k++ {
		l := rand.Intn(n)
		r := rand.Intn(n)
		if l < r {
			l, r = r, l
		}
		x := rand.Intn(lim)
		_ = pt.RangeLB(l, r, x)
	}
}
