package geometry3d

import (
	f "github.com/kelvinlau/go/floats"
	g2 "github.com/kelvinlau/go/geometry2d"
	"testing"
)

func TestIntersectionPointPlaneSeg(t *testing.T) {
	test := func(e Plane, l LineSeg, ok bool, ip Point) {
		if ok1, ip1 := IntersectionPointPlaneSeg(e, l); ok != ok1 || ok && !Coinside(ip, ip1) {
			t.Errorf("Expected %v %v, got %v %v.", ok, ip, ok1, ip1)
		}
	}

	e := Plane{Vector{1, 0, 0}, Point{0, 0, 0}}
	test(e, LineSeg{Point{1, 1, 1}, Point{-1, -1, -1}}, true, Point{0, 0, 0})
	test(e, LineSeg{Point{1, 1, 1}, Point{2, 2, 2}}, false, Point{})
	test(e, LineSeg{Point{1, 1, 1}, Point{1, 2, 1}}, false, Point{})
}

func TestOnPlane(t *testing.T) {
	test := func(e Plane, p Point, o bool) {
		if g := OnPlane(e, p); g != o {
			t.Errorf("%v, %v: expected %v, got %v.", e, p, o, g)
		}
	}

	e := Plane{Vector{0, 0, 1}, Point{0, 0, 0}}
	test(e, Point{-1, -1, 0}, true)
	test(e, Point{0, 0, 0}, true)
	test(e, Point{2, 2, 2}, false)
	test(e, Point{1, 2, 1}, false)
}

func TestDistPlanePoint(t *testing.T) {
	test := func(e Plane, p Point, o float64) {
		if g := DistPlanePoint(e, p); f.Sign2(g, o) != 0 {
			t.Errorf("%v, %v: expected %v, got %v.", e, p, o, g)
		}
	}

	e := Plane{Vector{0, 0, 1}, Point{0, 0, 0}}
	test(e, Point{-1, -1, 0}, 0)
	test(e, Point{0, 0, 0}, 0)
	test(e, Point{2, 2, 2}, 2)
	test(e, Point{1, 2, 1}, 1)
}

func TestPedal(t *testing.T) {
	test := func(e Plane, p, q Point) {
		if g := Pedal(e, p); !Coinside(g, q) {
			t.Errorf("%v, %v: expected %v, got %v.", e, p, q, g)
		}
	}

	e := Plane{Vector{0, 0, 1}, Point{0, 0, 0}}
	test(e, Point{-1, -1, 0}, Point{-1, -1, 0})
	test(e, Point{0, 0, 0}, Point{0, 0, 0})
	test(e, Point{2, 2, 2}, Point{2, 2, 0})
	test(e, Point{1, 2, 1}, Point{1, 2, 0})
}

func TestMap2D(t *testing.T) {
	e1 := Plane{Vector{1, 2, 3}, Point{4, 5, 6}}
	e2 := Plane{Vector{2, 0, 0}, Point{2, 0, 3}}
	for _, e := range []Plane{e1, e2} {
		ps := []Point{{0, 0, 0}, {1, 0, 1}, {2, 2, 0}}
		qs := []g2.Point{}
		for _, p := range ps {
			qs = append(qs, Map2D(e, p))
		}
		pp := Point{ps[1].X + ps[2].X - ps[0].X, ps[1].Y + ps[2].Y - ps[0].Y, ps[1].Z + ps[2].Z - ps[0].Z}
		qq := Map2D(e, pp)
		gg := g2.Point{qs[1].X + qs[2].X - qs[0].X, qs[1].Y + qs[2].Y - qs[0].Y}
		if f.Sign(g2.Dist(qq, gg)) != 0 {
			t.Errorf("Expected %v, got %v.", qq, gg)
		}
	}
}
