package floats

import (
	"testing"
)

func TestFloats(t *testing.T) {
	a := 1.2
	b := 1.1
	test := func(c float64, e int) {
		if g := Sign2(a*b, c); g != e {
			t.Errorf("%f * %f <> %f: expected %d, got %d.", a, b, c, e, g)
		}
	}
	c := 1.32
	c1 := 1.31
	c2 := 1.33
	test(c, 0)
	test(c1, +1)
	test(c2, -1)
}
