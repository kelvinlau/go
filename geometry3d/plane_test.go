package geometry3d

import (
	f "github.com/kelvinlau/go/floats"
	g2 "github.com/kelvinlau/go/geometry2d"
	"testing"
)

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
