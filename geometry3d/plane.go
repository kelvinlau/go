package geometry3d

import (
	f "github.com/kelvinlau/go/floats"
	g2 "github.com/kelvinlau/go/geometry2d"
	"math"
)

// Plane is a plane in 3d space, with normal vector N, passes point A.
type Plane struct {
	N Vector
	A Point
}

// IntersectionPointPlaneSeg is the intersection point of a plane and a line
// segment.
func IntersectionPointPlaneSeg(e Plane, l LineSeg) (ok bool, ip Point) {
	v := LineSegVec(l)
	lhs := Dot(e.N, v)
	if f.Sign(lhs) == 0 {
		return
	}
	rhs := Dot(e.N, Vec(l.P, e.A))
	t := rhs / lhs
	if f.Sign(t) >= 0 && f.Sign2(t, 1) <= 0 {
		ok = true
		ip = Add(l.P, Mul(v, t))
		return
	}
	return
}

// OnPlane reports if p is on e.
func OnPlane(e Plane, p Point) bool {
	return f.Sign(Dot(e.N, Vec(e.A, p))) == 0
}

// DistPlanePoint returns the distant of e and p.
func DistPlanePoint(e Plane, p Point) float64 {
	return math.Abs(Dot(e.N, Vec(e.A, p))) / Len(e.N)
}

// Pedal returns the closest point on e to p.
func Pedal(e Plane, p Point) Point {
	t := Dot(e.N, Vec(e.A, p)) / Len2(e.N)
	return Add(p, Mul(e.N, t))
}

// Map2D projects a 3D point to a 2D point on e.
func Map2D(e Plane, p Point) g2.Point {
	u := Vector{1, 0, 0}
	if Zero(Cross(e.N, u)) {
		u = Vector{0, 1, 0}
	}
	xd := Norm(Cross(e.N, u))
	yd := Norm(Cross(e.N, xd))
	ap := Vec(e.A, p)
	return g2.Point{Dot(ap, xd), Dot(ap, yd)}
}
