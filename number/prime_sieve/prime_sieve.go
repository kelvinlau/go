package prime_sieve

const (
	Max = 1000001 // All numbers < Max all be pre-processed.
)

var (
	ps []int
	fp [Max]int
)

func init() {
	for i := 2; i < Max; i++ {
		if fp[i] == 0 {
			fp[i] = i
			ps = append(ps, i)
		}
		for j := 0; ps[j]*i < Max; j++ {
			fp[ps[j]*i] = ps[j]
			if i%ps[j] == 0 {
				break
			}
		}
	}
}

// IsPrime repots if n is a prime, assuming n < Max * Max.
func IsPrime(n int) bool {
	if n < Max {
		return fp[n] == n
	}
	for i := 0; i < len(ps) && ps[i]*ps[i] <= n; i++ {
		if n%ps[i] == 0 {
			return false
		}
	}
	return true
}

// Factor is a component in n = p[0]^k[0] * p[1]^k[1] * ...
type Factor struct {
	p, k int
}

// Factorize factorize n, assuming n < Max * Max.
func Factorize(n int) (fs []Factor) {
	drain := func(p int) {
		k := 0
		for n%p == 0 {
			n /= p
			k++
		}
		fs = append(fs, Factor{p, k})
	}

	for i := 0; n >= Max && ps[i]*ps[i] <= n; i++ {
		if n%ps[i] == 0 {
			drain(ps[i])
		}
	}
	if n < Max {
		for n > 1 {
			drain(fp[n])
		}
	}
	if n > 1 {
		drain(n)
	}
	return
}
