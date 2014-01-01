package geometry2d

import (
	. "github.com/kelvinlau/go/floats"
	"math"
	"math/rand"
)

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
	return Sign(Dist(p, c.Point) - c.R)
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
	my2 := Sqr(my)
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

// CircleCircleIntersectionArea returns the intersection area of 2 circles.
func CircleCircleIntersectionArea(a, b Circle) float64 {
	d := Dist(a.Point, b.Point)
	if Sign(d) <= 0 || d+a.R <= b.R || d+b.R <= a.R {
		return Sqr(math.Min(a.R, b.R)) * math.Pi
	}
	if d >= a.R+b.R {
		return 0
	}

	da := (Sqr(d) + Sqr(a.R) - Sqr(b.R)) / d / 2
	db := d - da
	return Sqr(a.R)*math.Acos(da/a.R) - da*math.Sqrt(Sqr(a.R)-Sqr(da)) + Sqr(b.R)*math.Acos(db/b.R) - db*math.Sqrt(Sqr(b.R)-Sqr(db))
}

// CircleTagents returns 2 points of tangency on circle c of point p.
func CircleTagents(c Circle, p Point) (a, b Point) {
	d := Sqr(c.X-p.X) + Sqr(c.Y-p.Y)
	para := Sqr(c.R) / d
	perp := c.R * math.Sqrt(d-Sqr(c.R)) / d
	a.X = c.X + (p.X-c.X)*para - (p.Y-c.Y)*perp
	a.Y = c.Y + (p.Y-c.Y)*para + (p.X-c.X)*perp
	b.X = c.X + (p.X-c.X)*para + (p.Y-c.Y)*perp
	b.Y = c.Y + (p.Y-c.Y)*para - (p.X-c.X)*perp
	return
}

// Circle2 returns the minimum circle that covers 2 points.
func Circle2(a, b Point) Circle {
	return Circle{MidPoint(a, b), Dist(a, b) / 2}
}

// Circle3 returns the minimum circle that covers 3 points.
func Circle3(a, b, c Point) Circle {
	if d := Circle2(a, b); RelationCirclePoint(d, c) <= 0 {
		return d
	}
	if d := Circle2(b, c); RelationCirclePoint(d, a) <= 0 {
		return d
	}
	if d := Circle2(c, a); RelationCirclePoint(d, b) <= 0 {
		return d
	}
	o := CircumCircleCenter(a, b, c)
	r := Dist(o, a)
	return Circle{o, r}
}

// MinCircleCover returns the minimum circle that covers n points.
func MinCircleCover(ps []Point) Circle {
	if len(ps) == 0 {
		return Circle{Point{0, 0}, 0}
	}
	if len(ps) == 1 {
		return Circle{ps[0], 0}
	}
	if len(ps) == 2 {
		return Circle2(ps[0], ps[1])
	}
	if len(ps) == 3 {
		return Circle3(ps[0], ps[1], ps[2])
	}
	for i := range ps {
		j := rand.Intn(i + 1)
		ps[i], ps[j] = ps[j], ps[i]
	}
	qs := [4]*Point{&ps[0], &ps[1], &ps[2], &ps[3]}
	c := Circle3(ps[0], ps[1], ps[2])
	for {
		b := &ps[0]
		for i := 1; i < len(ps); i++ {
			if Dist(ps[i], c.Point) > Dist(*b, c.Point) {
				b = &ps[i]
			}
		}
		if RelationCirclePoint(c, *b) <= 0 {
			return c
		}
		qs[3] = b
		for i := 0; i < 3; i++ {
			qs[i], qs[3] = qs[3], qs[i]
			c = Circle3(*qs[0], *qs[1], *qs[2])
			if RelationCirclePoint(c, *qs[3]) <= 0 {
				break
			}
		}
	}
}
