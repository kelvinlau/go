package geometry2d

import (
	f "github.com/kelvinlau/go/floats"
	"math"
	"testing"
)

func TestDistLineSegPoint(t *testing.T) {
	test := func(l LineSeg, p Point, e float64) {
		if g := DistLineSegPoint(l, p); f.Sign2(g, e) != 0 {
			t.Errorf("%v %v, expected %f, got %f.", l, p, e, g)
		}
	}
	a := Point{0, 0}
	b := Point{1, 0}
	c := Point{1, 1}
	d := Point{-1, 1}
	e := Point{0.5, 0}
	f := Point{0.5, 1}
	l := LineSeg{a, b}
	test(l, a, 0)
	test(l, c, 1)
	test(l, d, math.Sqrt2)
	test(l, e, 0)
	test(l, f, 1)
}
