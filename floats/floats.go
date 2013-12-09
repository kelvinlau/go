package floats

const (
	Eps = 1E-8
)

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

func Sign2(x, y float64) int {
	return Sign(x - y)
}

func Sqr(x float64) float64 {
	return x * x
}
