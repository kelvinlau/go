package geometry3d

import (
	f "github.com/kelvinlau/go/floats"
)

// Plane is a plane in 3d space, with normal vector N, passes point A.
type Plane struct {
	N Vector
	A Point
}

// IntersectionPointPlaneSeg is the intersection point of a plane and a line
// segment.
func IntersectionPointPlaneSeg(e Plane, l LineSeg) (ok bool, ip Point) {
	lhs := Dot(e.N, LineSegVec(l))
	if f.Sign(lhs) == 0 {
		return
	}
	v := Vec(l.P, e.A)
	rhs := Dot(e.N, v)
	t := rhs / lhs
	if f.Sign(t) >= 0 && f.Sign2(t, 1) <= 0 {
		ok = true
		ip = Add(l.P, Mul(v, t))
		return
	}
	return
}
