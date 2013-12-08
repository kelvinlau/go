package geometry2d

import (
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

// TrianglesIntersectionArea returns the intersection area of 2 triangles.
func TrianglesIntersectionArea(a, b Triangle) float64 {
	ps := []Point{}
	ts := []Triangle{a, b}
	for t := 0; t < 2; t++ {
		for i := 0; i < 3; i++ {
			if InsideTriangleExclusive(ts[t][0], ts[t][1], ts[t][2], ts[1-t][i]) {
				ps = append(ps, ts[1-t][i])
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			l0 := LineSeg{ts[0][i], ts[0][(i+1)%3]}
			l1 := LineSeg{ts[1][j], ts[1][(j+1)%3]}
			if ok, p := LineSegIntersectionPoint(l0, l1); ok {
				ps = append(ps, p)
			}
		}
	}
	if len(ps) < 3 {
		return 0.0
	}
	AngularSort(ps)
	return AreaPolygon(ps)
}
