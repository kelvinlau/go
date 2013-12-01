package number

// IsPrimitiveRoot reports if w is a primitive root of r.
func IsPrimitiveRoot(r, w int64) bool {
	n := r - 1
	if w <= 1 || ModularPower(w, n, r) != 1 {
		return false
	}
	for d := int64(2); d*d <= n; d++ {
		if n%d == 0 {
			if ModularPower(w, d, r) == 1 || ModularPower(w, n/d, r) == 1 {
				return false
			}
		}
	}
	return true
}

// PrimitiveRoot returns a primitive root of n.
func PrimitiveRoot(n int64) int64 {
	for x := int64(2); x < n; x++ {
		if IsPrimitiveRoot(x, n) {
			return x
		}
	}
	return 0
}
