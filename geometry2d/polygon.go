// Package geometry2d is a 2D geometry library.
package geometry2d

import (
	"container/list"
	. "github.com/kelvinlau/go/floats"
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
	if len(ps) == 0 {
		return
	}
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
	q := func(i int) Point {
		return ps[b[i]]
	}
	r := func(i int) Point {
		return ps[(b[i]+1)%len(ps)]
	}
	for i, p := range ps {
		if LessY(p, q(0)) {
			b[0] = i
		}
		if LessX(q(1), p) {
			b[1] = i
		}
		if LessY(q(2), p) {
			b[2] = i
		}
		if LessX(p, q(3)) {
			b[3] = i
		}
	}

	alpha := float64(0)
	for k := 0; k < len(ps)+5; k++ {
		bi := -1
		minGap := math.Inf(1)
		for i := 0; i < 4; i++ {
			gap := Fix(Angle(q(i), r(i)) - (alpha + float64(i)*math.Pi/2))
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

		a := ShadowLength(alpha+math.Pi/2, q(0), q(2))
		b := ShadowLength(alpha, q(1), q(3))
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
			if p := LineSegIntersectionPoint(l, m); p != nil {
				qs = append(qs, *p)
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

// Triangulate returns n-2 triangles, which are the triangulation of a polygon.
func Triangulate(ps []Point) (ts []Triangle) {
	if len(ps) < 3 {
		return
	}
	qs := list.New()
	for _, p := range ps {
		qs.PushBack(p)
	}
	a := qs.Front()
	b := a.Next()
	c := b.Next()
	for ; c != nil; a, b, c = b, b.Next(), b.Next().Next() {
		A, B, C := a.Value.(Point), b.Value.(Point), c.Value.(Point)
		if Sign(Cross(A, B, C)) <= 0 {
			continue
		}
		ok := true
		for d := qs.Front(); d != nil; d = d.Next() {
			if d == a {
				d = d.Next().Next()
				continue
			}
			D := d.Value.(Point)
			if InsideTriangleInclusive(A, B, C, D) {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		ts = append(ts, Triangle{A, B, C})
		qs.Remove(b)
		b = a
		if b != qs.Front() {
			b = b.Prev()
		}
	}
	return
}

// PolygonIntersectionArea returns the intersection area of 2 simple polygons.
func PolygonIntersectionArea(ps, qs []Point) float64 {
	ts1 := Triangulate(ps)
	ts2 := Triangulate(qs)
	area := 0.0
	for _, t1 := range ts1 {
		for _, t2 := range ts2 {
			area += TrianglesIntersectionArea(t1, t2)
		}
	}
	return area
}

// LineIntersectionsConvexHull returns the convex hull from all intersection
// points of given lines in O(n).
func LineIntersectionsConvexHull(ls []Line) []Point {
	for i := range ls {
		if LessX(ls[i].Q, ls[i].P) {
			ls[i].P, ls[i].Q = ls[i].Q, ls[i].P
		}
	}
	sort.Sort(LinesAngleComparator(ls))
	l := []Line{}
	for i, j := 0, 0; i < len(ls); i = j {
		for j < len(ls) && Parallel(ls[i], ls[j]) {
			j++
		}
		l = append(l, ls[i])
		if j-i > 1 {
			l = append(l, ls[j-1])
		}
	}
	if len(l) < 2 {
		return []Point{}
	}
	l = append(l, l[0], l[1])

	ps := []Point{}
	for i, j := 0, 0; i < len(l); i++ {
		for j < len(l) && Parallel(l[i], l[j]) {
			j++
		}
		for k := j; k < len(l) && Parallel(l[j], l[k]); k++ {
			ps = append(ps, IntersectionPoint(l[i], l[k]))
		}
	}
	return ConvexHull(ps)
}
