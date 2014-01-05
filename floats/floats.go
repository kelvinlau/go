// Package floats implements basic floating point number util functions.
package floats

const (
	Eps = 1E-8 // Absolute error.
)

// Sign returns 0 for 0, +1 for positive number, -1 or negative number.
func Sign(x float64) int {
	switch {
	case x < -Eps:
		return -1
	case x > +Eps:
		return +1
	default:
		return 0
	}
}

// Sign2 return the sign of x-y.
func Sign2(x, y float64) int {
	return Sign(x - y)
}

// Eq returns true iff x == y within error.
func Eq(x, y float64) bool {
	return Sign(x-y) == 0
}

// Sqr returns the square of x.
func Sqr(x float64) float64 {
	return x * x
}
