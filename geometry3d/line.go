package geometry3d

import (
	f "github.com/kelvinlau/go/floats"
	"math"
)

// LineSeg is a line segment connect 2 points.
type LineSeg struct {
	P, Q Point
}

// Line is an infinite line passes through 2 points.
type Line LineSeg

// LineVec is the vector of line.
func LineVec(l Line) Vector {
	return Vec(l.P, l.Q)
}

// LineSegVec is the vector of line segment.
func LineSegVec(l LineSeg) Vector {
	return Vec(l.P, l.Q)
}

// OnLineSeg reports if a is on l.
func OnLineSeg(l LineSeg, a Point) bool {
	ap := Vec(a, l.P)
	aq := Vec(a, l.Q)
	return Zero(Cross(ap, aq)) && f.Sign(Dot(ap, aq)) <= 0
}

// Side returns the relationship of a, b based on l:
// +1: same plane, same side
// -1: same plane, opposite side
// 0:  otherwise
func Side(l LineSeg, a, b Point) int {
	v := LineSegVec(l)
	pa := Vec(l.P, a)
	pb := Vec(l.P, b)
	return f.Sign(Dot(Cross(v, pa), Cross(v, pb)))
}

// Touched reports if l1, l2 share a common point.
func Touched(l1, l2 LineSeg) bool {
	v1 := LineSegVec(l1)
	v2 := LineSegVec(l2)
	if Zero(Cross(v1, v2)) {
		return OnLineSeg(l1, l2.P) || OnLineSeg(l1, l2.Q) || OnLineSeg(l2, l1.P) || OnLineSeg(l2, l1.Q)
	} else {
		return Side(l1, l2.P, l2.Q) <= 0 && Side(l2, l1.P, l1.Q) <= 0
	}
}

// ClosestLinePoint returns the closest point on l to a.
func ClosestLinePoint(l Line, a Point) Point {
	v := LineVec(l)
	ap := Vec(a, l.P)
	t := Dot(v, ap) / Len2(v)
	return Add(l.P, Mul(v, t))
}

// ClosestLineSegPoint returns the closest point on l to a.
func ClosestLineSegPoint(l LineSeg, a Point) Point {
	v := LineSegVec(l)
	pa := Vec(l.P, a)
	t := math.Max(0, math.Min(1, Dot(v, pa)/Len2(v)))
	return Add(l.P, Mul(v, t))
}

// DistLineLine is the distant of 2 lines.
func DistLineLine(l1, l2 Line) float64 {
	v1 := LineVec(l1)
	v2 := LineVec(l2)
	if Zero(v1) || Zero(v2) {
		return 0
	}
	v3 := Vec(l1.P, l2.P)
	v4 := Vec(l1.P, l2.Q)
	v := Cross(v1, v2)
	if Zero(v) {
		return Len(Cross(v3, v4)) / Len(v2)
	}
	return math.Abs(Dot(v3, v)) / Len(v)
}

// ClosestApproach is like DistLineLine, but also returns the points of closest
// approach.
func ClosestApproach(l1, l2 Line) (d float64, p1, p2 Point) {
	v1 := LineVec(l1)
	v2 := LineVec(l2)
	v3 := Vec(l2.P, l1.P)
	s := Dot(v1, v2)
	t := Dot(v3, v2)
	den := Len2(v1)*Len2(v2) - f.Sqr(s)
	num := t*s - Len2(v2)*Dot(v3, v1)
	if f.Sign(den) == 0 {
		p1 = l1.P
		p2 = Add(l2.P, Mul(v2, t/Len2(v2)))
		if f.Sign(s) == 0 {
			p2 = p1
		}
	} else {
		num /= den
		p1 = Add(l1.P, Mul(v1, num))
		p2 = Add(l2.P, Mul(v2, (t+s*num)/Len2(v2)))
	}
	d = Len(Vec(p1, p2))
	return
}
