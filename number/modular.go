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
