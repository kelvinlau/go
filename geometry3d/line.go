package geometry3d

import (
	f "github.com/kelvinlau/go/floats"
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

// Pedal returns the meeting point where a line passes a with l at right angle.
func Pedal(l Line, a Point) Point {
	v := LineVec(l)
	ap := Vec(a, l.P)
	t := Dot(v, ap) / Len2(v)
	return Add(l.P, Mul(v, t))
}
