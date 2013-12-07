package geometry2d

import "math"

// Line is a line going through 2 points.
type Line struct {
	P, Q Point
}

// LineSeg is a line segment going through 2 points.
type LineSeg Line

// Ray is ray going from P through Q.
type Ray Line

// Parallel reports whether u and v are parallel.
func Parallel(u, v Line) bool {
	return Sign((u.P.X-u.Q.X)*(v.P.Y-v.Q.Y)-(v.P.X-v.Q.X)*(u.P.Y-u.Q.Y)) == 0
}

// Side returns 1 if p and q are on the same side of m; 0 if at least 1 of p, q
// touch m; otherwise -1.
func Side(m Line, p, q Point) int {
	return Sign(Cross(m.P, m.Q, p)) * Sign(Cross(m.P, m.Q, q))
}

// OnLine returns true iff p is on the line l.
func OnLine(l Line, p Point) bool {
	return Sign(Cross(l.P, l.Q, p)) == 0
}

// Coinside returns true if line u, v are the same line.
func Coinside(u, v Line) bool {
	return OnLine(u, v.P) && OnLine(u, v.Q)
}

// Intersected returns true if line segment u, v are intersected at 1 point,
// including the end points.
func Intersected(u, v LineSeg) bool {
	return !Parallel(Line(u), Line(v)) && Side(Line(u), v.P, v.Q) <= 0 && Side(Line(v), u.P, u.Q) <= 0
}

// IntersectedExclusive returns true if line segment u, v are intersected at 1
// point, not including the end points.
func IntersectedExclusive(u, v LineSeg) bool {
	return !Parallel(Line(u), Line(v)) && Side(Line(u), v.P, v.Q) < 0 && Side(Line(v), u.P, u.Q) < 0
}

// IntersectionPoint returns the intersection point of u and v, assuming u and v
// are not parallel to each other.
func IntersectionPoint(u, v Line) Point {
	n := (u.P.Y-v.P.Y)*(v.Q.X-v.P.X) - (u.P.X-v.P.X)*(v.Q.Y-v.P.Y)
	d := (u.Q.X-u.P.X)*(v.Q.Y-v.P.Y) - (u.Q.Y-u.P.Y)*(v.Q.X-v.P.X)
	r := n / d
	return Point{u.P.X + r*(u.Q.X-u.P.X), u.P.Y + r*(u.Q.Y-u.P.Y)}
}

// LineSegIntersectionPoint returns true and the intersection point of u and v,
// otherwise false and an arbitrary point.
func LineSegIntersectionPoint(u, v LineSeg) (ok bool, p Point) {
	if Parallel(Line(u), Line(v)) {
		ok = false
		return
	}
	p = IntersectionPoint(Line(u), Line(v))
	ok = OnLineSeg(u, p) && OnLineSeg(v, p)
	return
}

// OnLineSeg returns true iff p is on the line segment l, including the end
// points.
func OnLineSeg(l LineSeg, p Point) bool {
	return OnLine(Line(l), p) && Sign(Dot(p, l.P, l.Q)) <= 0
}

// OnLineSegExclusive returns true iff p is on the line segment l, excluding the
// end points.
func OnLineSegExclusive(l LineSeg, p Point) bool {
	return OnLine(Line(l), p) && Sign(Dot(p, l.P, l.Q)) < 0
}

// DistLinePoint returns the shortest distant between a point and a line.
func DistLinePoint(l Line, p Point) float64 {
	return math.Abs(Cross(l.P, l.Q, p)) / Dist(l.P, l.Q)
}

// DistLineSegPoint returns the shortest distant between a point and a line
// segment.
func DistLineSegPoint(l LineSeg, p Point) float64 {
	if OnLineSeg(l, p) {
		return 0
	}
	if Sign(Dot(l.P, l.Q, p)) <= 0 || Sign(Dot(l.Q, l.P, p)) <= 0 {
		return math.Min(Dist(l.P, p), Dist(l.Q, p))
	}
	return math.Abs(Cross(l.P, l.Q, p)) / Dist(l.P, l.Q)
}

// IntersectedLineSegRay return true, intersection point if u and v are
// intersected, otherwise false, (0, 0).
func IntersectedLineSegRay(u LineSeg, v Ray) (bool, Point) {
	if Parallel(Line(u), Line(v)) {
		return false, Point{}
	}
	p := IntersectionPoint(Line(u), Line(v))
	if OnLineSeg(u, p) && (OnLineSeg(LineSeg(v), p) || OnLineSeg(LineSeg{v.P, p}, v.Q)) {
		return true, p
	}
	return false, Point{}
}
