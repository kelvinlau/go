package geometry2d

import (
	. "github.com/kelvinlau/go/floats"
	"math"
)

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

// LineSegIntersectionPoint returns the intersection point of u and v,
// or nil if not intersected.
func LineSegIntersectionPoint(u, v LineSeg) *Point {
	if Parallel(Line(u), Line(v)) {
		return nil
	}
	p := IntersectionPoint(Line(u), Line(v))
	if OnLineSeg(u, p) && OnLineSeg(v, p) {
		return &p
	}
	return nil
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

// Perpendicular returns a line that goes through a point and meets l at a right
// angle.
func Perpendicular(l Line, p Point) Line {
	return Line{p, Point{p.X + l.P.Y - l.Q.Y, p.Y + l.Q.X - l.P.X}}
}

// Pedal return a point where l meets its perpendicular goes through p.
func Pedal(l Line, p Point) Point {
	return IntersectionPoint(l, Perpendicular(l, p))
}

// Mirror returns a point that reflect p on line l.
func Mirror(l Line, p Point) Point {
	q := Pedal(l, p)
	return Point{q.X*2 - p.X, q.Y*2 - p.Y}
}

// PerpendicularBisector returns a line perpendicular to ab and split them
// equally.
func PerpendicularBisector(a, b Point) Line {
	return Perpendicular(Line{a, b}, MidPoint(a, b))
}

// Translate returns the new location where l will translate with distance e and
// direction s.
func Translate(l Line, e float64, s int) Line {
	d := Dist(l.P, l.Q)
	x := l.P.Y - l.Q.Y
	y := l.Q.X - l.P.X
	x *= float64(s) * e / d
	y *= float64(s) * e / d
	l.P.X += x
	l.P.Y += y
	l.Q.X += x
	l.Q.Y += y
	return l
}

// LinesAngleComparator is a sort.Interface the sorts lines by angle.
type LinesAngleComparator []Line

func (ls LinesAngleComparator) Len() int { return len(ls) }
func (ls LinesAngleComparator) Less(i, j int) bool {
	u, v := ls[i], ls[j]
	c := Sign((u.P.X-u.Q.X)*(v.P.Y-v.Q.Y) - (v.P.X-v.Q.X)*(u.P.Y-u.Q.Y))
	return c < 0 || c == 0 && Sign(Cross(u.P, u.Q, v.P)) < 0
}
func (ls LinesAngleComparator) Swap(i, j int) { ls[i], ls[j] = ls[j], ls[i] }
