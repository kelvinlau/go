package number

import (
	"math"
	"sort"
)

// Returns x ^ y % r.
func ModularPower(x, y, r int64) int64 {
	z := int64(1)
	for y > 0 {
		if y&1 != 0 {
			z = z * x % r
		}
		x = x * x % r
		y >>= 1
	}
	return z
}

// Returns y, such that x * y % r == 1
func ModularInvert(x, r int64) int64 {
	return ModularPower(x, r-2, r)
}

// GCD returns the greatest common divisor of x and y.
func GCD(x, y int64) int64 {
	if y == 0 {
		return x
	}
	return GCD(y, x%y)
}

// LCM returns the least common multiplier of x and y.
func LCM(x, y int64) int64 {
	return x / GCD(x, y) * y
}

// ExGCD returns a tuple (a, b, g) such that m * a + n * b == g and
// g == GCD(m, n).
func ExGCD(m, n int64) (a, b, g int64) {
	if n == 0 {
		return 1, 0, m
	}
	b, a, g = ExGCD(n, m%n)
	b -= m / n * a
	return
}

// ModularSystem returns x such that x % ms[i] == rs[i] for all i.
func ModularSystem(ms, rs []int64) int64 {
	var m, r int64 = 1, 0
	for i := 0; i < len(ms) && r != -1; i++ {
		r = modularSystem(m, r, ms[i], rs[i])
		m = LCM(m, ms[i])
	}
	return r
}

// modularSystem returns x such that x % m == a && x % n == b.
func modularSystem(m, a, n, b int64) int64 {
	k, _, g := ExGCD(m, n)
	if (a-b)%g != 0 {
		return -1
	}
	k *= (b - a) / g
	k = mod(k, n/g)
	return mod(k*m+a, m/g*n)
}

func mod(x, y int64) int64 {
	return (x%y + y) % y
}

type pair struct {
	p int64
	j int
}
type pairs []pair

func (a pairs) Len() int           { return len(a) }
func (a pairs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a pairs) Less(i, j int) bool { return a[i].p < a[j].p }

// ModularLog returns x such that a ^ x % m == b.
func ModularLog(a, b, m int64) int64 {
	s := int(math.Ceil(math.Sqrt(float64(m))))
	ts := make(pairs, s)
	for j, p := 0, int64(1); j < s; j, p = j+1, p*a%m {
		ts[j] = pair{p, j}
	}
	sort.Sort(ts)

	c := ModularInvert(ModularPower(a, int64(s), m), m) // c = a^(-m).
	for i := 0; i < s; i++ {
		k := sort.Search(s, func(k int) bool { return ts[k].p >= b })
		if k < s && ts[k].p == b {
			return int64(i*s + ts[k].j)
		}
		b = b * c % m
	}
	return -1
}
