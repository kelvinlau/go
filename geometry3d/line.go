package geometry3d

import (
	f "github.com/kelvinlau/go/floats"
)

// LineSeg is a line segment connect 2 points.
type LineSeg struct {
	P, Q Point
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
	v := Vec(l.P, l.Q)
	pa := Vec(l.P, a)
	pb := Vec(l.P, b)
	return f.Sign(Dot(Cross(v, pa), Cross(v, pb)))
}
