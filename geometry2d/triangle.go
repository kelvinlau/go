package geometry2d

import (
	"container/list"
	"math"
)

// Triangle is a polygon with 3 points.
type Triangle [3]Point

// Calculates the area of a triangle given the lengths of sides.
func AreaHeron(a, b, c float64) float64 {
	s := (a + b + c) / 2
	if a > s || b > s || c > s {
		return -1
	}
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

// Calculates the area of a triangle given 3 points.
func AreaTriangle(a, b, c Point) float64 {
	return math.Abs(Cross(a, b, c)) / 2
}

// Sharp returns true if ABC is a sharp triangle
func Sharp(a, b, c Point) bool {
	return Sign(Dot(a, b, c)) > 0 && Sign(Dot(b, a, c)) > 0 && Sign(Dot(c, a, b)) > 0
}

// Perpencenter returns the meeting point of 3 altitudes of the triangle abc.
func Perpencenter(a, b, c Point) Point {
	u := Perpendicular(Line{b, c}, a)
	v := Perpendicular(Line{a, c}, b)
	return IntersectionPoint(u, v)
}

// InsideTriangleExclusive returns true iff p is inside the triangle abc, not
// including the border.
func InsideTriangleExclusive(a, b, c, p Point) bool {
	return Sign(Cross(a, b, p)) > 0 && Sign(Cross(b, c, p)) > 0 && Sign(Cross(c, a, p)) > 0
}

// InsideTriangleInclusive returns true iff p is inside the triangle abc,
// including the border.
func InsideTriangleInclusive(a, b, c, p Point) bool {
	return Sign(Cross(a, b, p)) >= 0 && Sign(Cross(b, c, p)) >= 0 && Sign(Cross(c, a, p)) >= 0
}

// Trianglulate returns n-2 triangles, which are the triangulation of a polygon.
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
