package geometry2d

import "math"

// Circle with center and radius.
type Circle struct {
	Point
	R float64
}

// InscribedCircleCenter returns the center of the max circle that fits into
// the triangle abc.
func InscribedCircleCenter(A, B, C Point) Point {
	a := Dist(B, C)
	b := Dist(C, A)
	c := Dist(A, B)
	p := a + b + c
	return Point{(a*A.X + b*B.X + c*C.X) / p, (a*A.Y + b*B.Y + c*C.Y) / p}
}

// CircumCircleCenter returns the center of the min circle that contains the
// triangle abc.
func CircumCircleCenter(A, B, C Point) Point {
	a1 := B.X - A.X
	b1 := B.Y - A.Y
	c1 := (Sqr(a1) + Sqr(b1)) / 2
	a2 := C.X - A.X
	b2 := C.Y - A.Y
	c2 := (Sqr(a2) + Sqr(b2)) / 2
	d := a1*b2 - a2*b1
	return Point{A.X + (c1*b2-c2*b1)/d, A.Y + (a1*c2-a2*c1)/d}
}

// RelationCirclePoint returns 0 if p is on c, -1 if inside, +1 if outside.
func RelationCirclePoint(c Circle, p Point) int {
	return Sign(c.R - Dist(p, c.Point))
}

// RelationCircleLine returns 0 if l is tangent to c, -1 if intersected, +1 if
// outside.
func RelationCircleLine(c Circle, l Line) int {
	return Sign(DistLinePoint(l, c.Point) - c.R)
}

// RelationCircleLineSeg returns 0 if l just touches c, -1 if inside of
// intersected, +1 if outside.
func RelationCircleLineSeg(c Circle, l LineSeg) int {
	return Sign(DistLineSegPoint(l, c.Point) - c.R)
}

// HasCommonPointCircleTriangle returns true iff c and t share common points.
func HasCommonPointCircleTriangle(c Circle, t Triangle) bool {
	for _, p := range t {
		if RelationCirclePoint(c, p) <= 0 {
			return true
		}
	}
	for i, _ := range t {
		j := (i + 1) % 3
		l := LineSeg{t[i], t[j]}
		if RelationCircleLineSeg(c, l) <= 0 {
			return true
		}
	}
	return InsidePolygon(t[:], c.Point)
}

// IntersectionPointCircleLine returns the intersection points of c and l.
// Results can be of size 0, 1, or 2.
func IntersectionPointCircleLine(c Circle, l Line) []Point {
	a := l.P
	b := l.Q
	dx := b.X - a.X
	dy := b.Y - a.Y
	sdr := Sqr(dx) + Sqr(dy)
	a.X -= c.X
	a.Y -= c.Y
	b.X -= c.X
	b.Y -= c.Y
	d := a.X*b.Y - b.X*a.Y
	disc := Sqr(c.R)*sdr - Sqr(d)
	if Sign(disc) < 0 {
		return []Point{}
	}
	if Sign(disc) == 0 {
		disc = 0
	} else {
		disc = math.Sqrt(disc)
	}
	s := 0.0
	if dy > 0 {
		s = +1
	} else {
		s = -1
	}
	x := disc * dx * s
	y := disc * math.Abs(dy)
	a.X = (+d*dy+x)/sdr + c.X
	b.X = (+d*dy-x)/sdr + c.X
	a.Y = (-d*dx+y)/sdr + c.Y
	b.Y = (-d*dx-y)/sdr + c.Y
	if Sign(disc) > 0 {
		return []Point{a, b}
	} else {
		return []Point{a}
	}
}

// IntersectionPointCircleCircle returns the intersection points of two circles,
// Results can be of size 0, 1, or 2.
func IntersectionPointCircleCircle(c1, c2 Circle) []Point {
	mx := c2.X - c1.X
	sx := c2.X + c1.X
	mx2 := Sqr(mx)
	my := c2.Y - c1.Y
	sy := c2.Y + c1.Y
	my2 := Sqr(mx)
	sq := mx2 + my2
	d := -(sq - Sqr(c1.R-c2.R)) * (sq - Sqr(c1.R+c2.R))
	if Sign(sq) == 0 || Sign(d) < 0 {
		return []Point{}
	}
	if Sign(d) == 0 {
		d = 0
	} else {
		d = math.Sqrt(d)
	}
	x := mx*((c1.R+c2.R)*(c1.R-c2.R)+mx*sx) + sx*my2
	y := my*((c1.R+c2.R)*(c1.R-c2.R)+my*sy) + sy*mx2
	dx := mx * d
	dy := my * d
	sq *= 2
	a := Point{(x + dy) / sq, (y - dx) / sq}
	b := Point{(x - dy) / sq, (y + dx) / sq}
	if Sign(d) > 0 {
		return []Point{a, b}
	} else {
		return []Point{a}
	}
}
