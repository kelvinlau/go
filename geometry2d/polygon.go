package geometry2d

import (
	"math"
)

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
