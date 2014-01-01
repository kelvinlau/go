package geometry2d

import (
	. "github.com/kelvinlau/go/floats"
	"math"
	"testing"
)

func TestIntersectionPointCircleLine(t *testing.T) {
	test := func(c Circle, l Line, e int) {
		if g := IntersectionPointCircleLine(c, l); len(g) != e {
			t.Errorf("Expected %d intersection points, got %d.", e, len(g))
			for _, p := range g {
				if !OnLine(l, p) || RelationCirclePoint(c, p) != 0 {
					t.Errorf("Invalid intersection point %v of %v and %v.", p, c, l)
				}
			}
		}
	}
	c := Circle{Point{0, 0}, 1}
	l1 := Line{Point{0, 0}, Point{0, 1}}
	l2 := Line{Point{1, 0}, Point{1, 1}}
	l3 := Line{Point{2, 0}, Point{2, 1}}
	test(c, l1, 2)
	test(c, l2, 1)
	test(c, l3, 0)
}

func TestIntersectionPointCircleCircle(t *testing.T) {
	test := func(c1, c2 Circle, e int) {
		if g := IntersectionPointCircleCircle(c1, c2); len(g) != e {
			t.Errorf("%v, %v: expected %d intersection points, got %d.", c1, c2, e, len(g))
			for _, p := range g {
				if RelationCirclePoint(c1, p) != 0 || RelationCirclePoint(c2, p) != 0 {
					t.Errorf("Invalid intersection point %v of %v and %v.", p, c1, c2)
				}
			}
		}
	}
	c := Circle{Point{0, 0}, 1}
	c1 := Circle{Point{0, 0}, 2}
	c2 := Circle{Point{0, 1}, 1}
	c3 := Circle{Point{0, 2}, 1}
	c4 := Circle{Point{0, 3}, 1}
	test(c, c1, 0)
	test(c, c2, 2)
	test(c, c3, 1)
	test(c, c4, 0)
}

func TestCircleCircleIntersectionArea(t *testing.T) {
	c1 := Circle{Point{0, 0}, 1}
	c2 := Circle{Point{0, 1}, 1}
	e := math.Pi/3*2 - math.Sqrt(3)/2
	if g := CircleCircleIntersectionArea(c1, c2); Sign(g-e) != 0 {
		t.Errorf("Expected intersection area %f, got %f.", e, g)
	}
}

func TestMinCircleCover(t *testing.T) {
	test := func(ps []Point, e Circle) {
		c := MinCircleCover(ps)
		if Sign(c.R-e.R) != 0 || Sign(Dist(e.Point, c.Point)) != 0 {
			t.Errorf("Expected %v, got %v.", e, c)
		}
	}
	test([]Point{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}, Circle{Point{1, 1}, math.Sqrt2})
	test([]Point{
		{-1, 0},
		{+1, 0},
		{0, +2},
	}, Circle{Point{0, 0.75}, 1.25})
	test([]Point{
		{-1, -1},
		{+1, +1},
	}, Circle{Point{0, 0}, math.Sqrt2})
	test([]Point{
		{0, 0},
	}, Circle{Point{0, 0}, 0})
	test([]Point{}, Circle{Point{0, 0}, 0})
}

func TestCircleTagents(t *testing.T) {
	c := Circle{Point{0, 0}, 1}
	p := Point{2, 0}
	alpha := math.Pi / 3
	e1 := NextPoint(c.Point, -alpha, 1)
	e2 := NextPoint(c.Point, +alpha, 1)
	q1, q2 := CircleTagents(c, p)
	if LessY(q2, q1) {
		q1, q2 = q2, q1
	}
	if Sign(Dist(q1, e1)) > 0 || Sign(Dist(q2, e2)) > 0 {
		t.Errorf("Expected %v %v, got %v %v.", e1, e2, q1, q2)
	}
}
