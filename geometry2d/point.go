package geometry2d

import (
	"math"
)

// Point is a point on 2d plane.
type Point struct {
	X, Y float64
}

// Cross product of AB x BC.
func Cross(a, b, c Point) float64 {
	return (b.X-a.X)*(c.Y-b.Y) - (b.Y-a.Y)*(c.X-b.X)
}

// Dot returns dot product of PA * PB.
func Dot(p, a, b Point) float64 {
	return (a.X-p.X)*(b.X-p.X) + (a.Y-p.Y)*(b.Y-p.Y)
}

// Dist returns the distant between A and B.
func Dist(a, b Point) float64 {
	return math.Hypot(a.X-b.X, a.Y-b.Y)
}

// MidPoint returns mid-point of A and B.
func MidPoint(a, b Point) Point {
	return Point{(a.X + b.X) / 2, (a.Y + b.Y) / 2}
}

// NextPoint returns a point dist away from A to the angle of alpha.
func NextPoint(a Point, alpha, dist float64) Point {
	return Point{a.X + dist*math.Cos(alpha), a.Y + dist*math.Sin(alpha)}
}

// DeltaAngle returns delta angle of ABC.
func DeltaAngle(a, b, c Point) float64 {
	return math.Acos(Dot(b, a, c) / (Dist(a, b) * Dist(b, c)))
}

// Sharp returns true if ABC is a sharp triangle
func Sharp(a, b, c Point) bool {
	return Sign(Dot(a, b, c)) > 0 && Sign(Dot(b, a, c)) > 0 && Sign(Dot(c, a, b)) > 0
}

// Fix returns a angle capped in [0, 2*PI)
func Fix(a float64) float64 {
	if Sign(a) < 0 {
		a += 2 * math.Pi
	}
	if Sign(a-2*math.Pi) >= 0 {
		a -= 2 * math.Pi
	}
	return a
}

// Angle returns the angle of B base on A.
func Angle(a, b Point) float64 {
	return Fix(math.Atan2(b.Y-a.Y, b.X-a.X))
}

// Rotate returns the destination where point A rotate alpha around B.
func Rotate(a, b Point, alpha float64) Point {
	a.X -= b.X
	a.Y -= b.Y
	c := math.Cos(alpha)
	s := math.Sin(alpha)
	b.X += a.X*c - a.Y*s
	b.Y += a.X*s + a.Y*c
	return b
}

// ShadowLength returns the length of shadow projected onto the give angle.
func ShadowLength(alpha float64, a, b Point) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	c := math.Cos(alpha)
	s := math.Sin(alpha)
	return math.Abs(dx*c + dy*s)
}
