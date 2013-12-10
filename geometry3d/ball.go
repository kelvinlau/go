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

// GreatCircle is a function that given the latitude and longitude of two points
// in degrees, calculates the distance over the sphere between them.
// Latitude is given in the range [-pi/2, pi/2] degrees,
// Longitude is given in the range [-pi,pi] degrees.
func GreatCircle(lat1, long1, lat2, long2 float64) float64 {
	return math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(long2-long1))
}
