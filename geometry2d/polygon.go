package geometry2d

import (
	"math"
	"sort"
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

type angularCmp struct {
	o Point
	p []Point
}

func (a *angularCmp) Less(i, j int) bool {
	c := Sign(Cross(a.o, a.p[i], a.p[j]))
	return c > 0 || c == 0 && Dist(a.o, a.p[i]) < Dist(a.o, a.p[j])
}

func (a *angularCmp) Len() int {
	return len(a.p)
}

func (a *angularCmp) Swap(i, j int) {
	a.p[i], a.p[j] = a.p[j], a.p[i]
}

// AngularSort sorts the points by angular in place.
func AngularSort(p []Point) {
	o := p[0]
	for i := 1; i < len(p); i++ {
		if p[i].X < o.X || p[i].X == o.X && p[i].Y < o.Y {
			o = p[i]
		}
	}
	acmp := &angularCmp{o, p}
	sort.Sort(acmp)
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
