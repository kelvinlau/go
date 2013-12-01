package number

// IsPrime reports if n is a prime.
func IsPrime(n int64) bool {
	if n == 2 {
		return true
	}
	if n <= 1 || n%2 == 0 {
		return false
	}
	for p := int64(3); p*p <= n; p += 2 {
		if n%p == 0 {
			return false
		}
	}
	return true
}
