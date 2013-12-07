package geometry2d

import "math"

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
