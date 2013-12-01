package number

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
