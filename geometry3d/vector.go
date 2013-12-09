package geometry3d

import (
	f "github.com/kelvinlau/go/floats"
	"math"
)

// Vector is a direction on 3d space.
type Vector Point

// Vec is the Vector from a to b.
func Vec(a, b Point) Vector {
	return Sub(Vector(b), Vector(a))
}

// Len is |a|.
func Len(a Vector) float64 {
	return math.Sqrt(Len2(a))
}

// Len2 is |a| * |a|.
func Len2(a Vector) float64 {
	return Dot(a, a)
}

// Zeros is (|a| == 0).
func Zero(a Vector) bool {
	return f.Sign(a.X) == 0 && f.Sign(a.Y) == 0 && f.Sign(a.Z) == 0
}

// Add is a + b.
func Add(a, b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Sub is a - b.
func Sub(a, b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Mul is a * k.
func Mul(a Vector, k float64) Vector {
	return Vector{a.X * k, a.Y * k, a.Z * k}
}

// Div is a / k.
func Div(a Vector, k float64) Vector {
	return Vector{a.X / k, a.Y / k, a.Z / k}
}

// Norm is a / |a|.
func Norm(a Vector) Vector {
	return Div(a, Len(a))
}

// Dot is a . b (dot product).
func Dot(a, b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross is a x b (cross product).
func Cross(a, b Vector) Vector {
	return Vector{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X}
}
