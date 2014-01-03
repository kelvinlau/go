// Package geometry3d is a 3D geometry library.
package geometry3d

import (
	. "github.com/kelvinlau/go/floats"
)

// Point is a point on 3d space.
type Point struct {
	X, Y, Z float64
}

// Coinside returns true iff p equals q.
func Coinside(p, q Point) bool {
	return Eq(p.X, q.X) && Eq(p.Y, q.Y) && Eq(p.Z, q.Z)
}
