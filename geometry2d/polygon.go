package geometry2d

import (
	"math"
	"sort"
)

// Calculates the area of a polygon.
func AreaPolygon(a []Point) float64 {
	area := float64(0)
	n := len(a)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += a[i].X*a[j].Y - a[i].Y*a[j].X
	}
	return area / 2
}

// Calculates the centroid of a polygon.
func Centroid(a []Point) Point {
	c := Point{}
	area := float64(0)
	n := len(a)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += a[i].X*a[j].Y - a[i].Y*a[j].X
		c.X += (a[i].X + a[j].X) * (a[i].X*a[j].Y - a[i].Y*a[j].X)
		c.Y += (a[i].Y + a[j].Y) * (a[i].X*a[j].Y - a[i].Y*a[j].X)
	}
	area = math.Abs(area) / 2
	c.X /= 6 * area
	c.Y /= 6 * area
	return c
}

// Find convex hull.
func ConvexHull(ps []Point) (qs []Point) {
	AngularSort(ps)
	for _, p := range ps {
		for len(qs) >= 2 && Sign(Cross(qs[len(qs)-2], qs[len(qs)-1], p)) <= 0 {
			qs = qs[:len(qs)-1]
		}
		qs = append(qs, p)
	}
	return
}

// RotateCalipers returns the min covering rectangle area & perimeter of a
// convex hull, using rotating calipers algorithm, O(n).
func RotateCalipers(ps []Point) (area, peri float64) {
	area, peri = math.Inf(1), math.Inf(1)
	b := [4]int{}
	q := func(i int) *Point {
		return &ps[b[i]]
	}
	r := func(i int) *Point {
		return &ps[(b[i]+1)%len(ps)]
	}
	for i, p := range ps {
		if q(0).Y > p.Y || q(0).Y == p.Y && q(0).X > p.X {
			b[0] = i
		}
		if q(1).X < p.X || q(1).X == p.X && q(1).Y > p.Y {
			b[1] = i
		}
		if q(2).Y < p.Y || q(2).Y == p.Y && q(2).X < p.X {
			b[2] = i
		}
		if q(3).X > p.X || q(3).X == p.X && q(3).Y < p.Y {
			b[3] = i
		}
	}

	alpha := float64(0)
	for k := 0; k < len(ps)+5; k++ {
		bi := -1
		minGap := math.Inf(1)
		for i := 0; i < 4; i++ {
			gap := Fix(Angle(*q(i), *r(i)) - (alpha + float64(i)*math.Pi/2))
			if gap < minGap {
				minGap = gap
				bi = i
			}
		}
		b[bi]++
		if b[bi] == len(ps) {
			b[bi] = 0
		}
		alpha = Fix(alpha + minGap)

		a := ShadowLength(alpha+math.Pi/2, *q(0), *q(2))
		b := ShadowLength(alpha, *q(1), *q(3))
		area = math.Min(area, a*b)
		peri = math.Min(peri, a+a+b+b)
	}
	return
}

// InsidePolygon returns true iff a strickly inside polygon ps.
func InsidePolygon(ps []Point, a Point) bool {
	sum := 0.0
	for i := 0; i < len(ps); i++ {
		j := (i + 1) % len(ps)
		if OnLineSeg(LineSeg{ps[i], ps[j]}, a) {
			return false
		}
		angle := math.Acos(Dot(a, ps[i], ps[j]) / Dist(a, ps[i]) / Dist(a, ps[j]))
		sum += float64(Sign(Cross(a, ps[i], ps[j]))) * angle
	}
	return Sign(sum) > 0
}

// LineSegInsidePolygon returns true iff l strickly inside ps.
func LineSegInsidePolygon(ps []Point, l LineSeg) bool {
	for i := 0; i < len(ps); i++ {
		j := (i + 1) % len(ps)
		m := LineSeg{ps[i], ps[j]}
		if OnLineSegExclusive(l, ps[i]) {
			return false
		}
		if IntersectedExclusive(l, m) {
			return false
		}
	}
	return InsidePolygon(ps, MidPoint(l.P, l.Q))
}

// IntersectedConvexLineSeg returns true iff l intersect with convex hull ps.
func IntersectedConvexLineSeg(ps []Point, l LineSeg) bool {
	if len(ps) < 3 {
		return false
	}
	qs := []Point{l.P, l.Q}
	for i := 0; i < len(ps); i++ {
		if OnLineSeg(l, ps[i]) {
			qs = append(qs, ps[i])
		} else {
			j := (i + 1) % len(ps)
			m := LineSeg{ps[i], ps[j]}
			if ok, p := LineSegIntersectionPoint(l, m); ok {
				qs = append(qs, p)
			}
		}
	}

	sort.Sort(ByX{qs})
	for i := 0; i+1 < len(qs); i++ {
		if InsidePolygon(ps, MidPoint(qs[i], qs[i+1])) {
			return true
		}
	}
	return false
}

// CutArea returns the area of the part of convex hull ps on the positive side
// of l.
func CutArea(ps []Point, l Line) float64 {
	i1, i2 := -1, -1
	for i := 0; i < len(ps); i++ {
		j := (i + 1) % len(ps)
		v := LineSeg{ps[i], ps[j]}
		if Parallel(Line(v), l) {
			continue
		}
		if cp := IntersectionPoint(Line(v), l); cp == ps[i] || OnLineSegExclusive(v, cp) {
			if i1 == -1 {
				i1 = i
			} else {
				i2 = i
			}
		}
	}
	qs := []Point{}
	for i := 0; i < len(ps); i++ {
		if Sign(Cross(l.P, l.Q, ps[i])) == 0 {
			qs = append(qs, ps[i])
		}
		if i == i1 {
			qs = append(qs, ps[i1])
		}
		if i == i2 {
			qs = append(qs, ps[i2])
		}
	}
	return AreaPolygon(qs)
}