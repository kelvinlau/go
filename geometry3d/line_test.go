package geometry3d

import (
	. "github.com/kelvinlau/go/floats"
	"testing"
)

func TestTouched(t *testing.T) {
	test := func(l1, l2 LineSeg, e bool) {
		if g := Touched(l1, l2); g != e {
			t.Errorf("%v, %v: expected %v, got %v.", l1, l2, e, g)
		}
	}
	l1 := LineSeg{Point{-1, 0, 0}, Point{1, 0, 0}}
	l2 := LineSeg{Point{0, -1, 0}, Point{0, 1, 0}}
	l3 := LineSeg{Point{1, 1, 1}, Point{2, 2, 2}}
	test(l1, l1, true)
	test(l1, l2, true)
	test(l1, l3, false)
	test(l2, l3, false)
}

func TestClosestLineSegPoint(t *testing.T) {
	l := LineSeg{Point{0, 0, 0}, Point{1, 0, 0}}
	p := Point{-1, -1, 1}
	e := l.P
	if g := ClosestLineSegPoint(l, p); !Coinside(g, e) {
		t.Errorf("Expected %v, got %v.", e, g)
	}
}

func TestClosestApproach(t *testing.T) {
	test := func(l1, l2 Line, e float64, p1, p2 Point) {
		d, g1, g2 := ClosestApproach(l1, l2)
		if !Eq(d, e) || !Coinside(p1, g1) || !Coinside(p2, g2) {
			t.Errorf("Expected (%v, %v, %v), got (%v, %v, %v).", e, p1, p2, d, g1, g2)
		}
		if g := DistLineLine(l1, l2); !Eq(g, e) {
			t.Errorf("DistLineLine(%v, %v) = %v, expected %v.", l1, l2, g, d)
		}
	}
	{
		l1 := Line{Point{1, 1, 0}, Point{2, 2, 0}}
		l2 := Line{Point{1, -1, 3}, Point{2, -2, 3}}
		e := 3.0
		p1 := Point{0, 0, 0}
		p2 := Point{0, 0, 3}
		test(l1, l2, e, p1, p2)
	}
	{
		l1 := Line{Point{1, 1, 0}, Point{2, 2, 0}}
		l2 := Line{Point{2, 2, 3}, Point{1, 1, 3}}
		e := 3.0
		p1 := Point{1, 1, 0}
		p2 := Point{1, 1, 3}
		test(l1, l2, e, p1, p2)
	}
	{
		l1 := Line{Point{1, 1, 0}, Point{1, 1, 0}}
		l2 := Line{Point{2, 2, 3}, Point{1, 1, 3}}
		e := 0.0
		p1 := Point{1, 1, 0}
		p2 := Point{1, 1, 0}
		test(l1, l2, e, p1, p2)
	}
}
