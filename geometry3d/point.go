// Package geometry3d is a 3D geometry library.
package geometry3d

import (
	f "github.com/kelvinlau/go/floats"
)

// Point is a point on 3d space.
type Point struct {
	X, Y, Z float64
}

// Coinside returns true iff p equals q.
func Coinside(p, q Point) bool {
	return f.Sign2(p.X, q.X) == 0 && f.Sign2(p.Y, q.Y) == 0 && f.Sign2(p.Z, q.Z) == 0
}
