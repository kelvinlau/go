package geometry3d

import (
	. "github.com/kelvinlau/go/floats"
	"math"
	"testing"
)

func TestIntersectionPointBallLine(t *testing.T) {
	test := func(b Ball, l Line, e []Point) {
		g := IntersectionPointBallLine(b, l)
		if len(g) != len(e) {
			t.Fatalf("Expected %v, got %v.", e, g)
		}
		for _, p := range e {
			found := false
			for _, q := range g {
				found = found || Coinside(p, q)
			}
			if !found {
				t.Fatalf("Expected %v, got %v.", e, g)
			}
		}
	}

	b := Ball{Point{0, 0, 0}, 1}
	test(b, Line{Point{0, 0, 0}, Point{0, 2, 0}}, []Point{Point{0, 1, 0}, Point{0, -1, 0}})
	test(b, Line{Point{1, 0, 0}, Point{1, 2, 0}}, []Point{Point{1, 0, 0}})
	test(b, Line{Point{2, 0, 0}, Point{1, 2, 0}}, []Point{})
}

func TestGreatCircle(t *testing.T) {
	if g, e := GreatCircle(math.Pi/2, 0, 0, math.Pi), math.Pi/2; !Eq(g, e) {
		t.Errorf("Expected %v, got %v.", e, g)
	}
}

func TestBall4(t *testing.T) {
	test := func(ps [4]Point, e *Ball) {
		g := Ball4(ps)
		if (g == nil) != (e == nil) || g != nil && (!Coinside(e.Point, g.Point) || !Eq(e.R, g.R)) {
			t.Fatalf("Expected %v, got %v.", e, g)
		}
	}
	test([4]Point{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{-1, 0, 0}},
		&Ball{Point{0, 0, 0}, 1})

	test([4]Point{
		{1, 0, 0},
		{0, 1, 0},
		{0, 2, 0},
		{-1, 0, 0}},
		nil)
}
