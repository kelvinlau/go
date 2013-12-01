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
