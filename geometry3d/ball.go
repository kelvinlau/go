package geometry3d

import (
	f "github.com/kelvinlau/go/floats"
	"math"
)

type Ball struct {
	Point
	R float64
}

// IntersectionPointBallLine returns intersection points of a ball and a line,
// results may be of size 0, 1, or 2.
func IntersectionPointBallLine(b Ball, l Line) []Point {
	d := ClosestLinePoint(l, b.Point)
	e := Len2(Vec(b.Point, d))
	if e > f.Sqr(b.R) {
		return []Point{}
	}
	lv := LineVec(l)
	t := math.Sqrt(f.Sqr(b.R)-e) / Len2(lv)
	if f.Sign(t) == 0 {
		return []Point{d}
	}
	v := Mul(lv, t)
	return []Point{Add(d, v), Sub(d, v)}
}
