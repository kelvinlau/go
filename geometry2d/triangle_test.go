package geometry2d

import (
	f "github.com/kelvinlau/go/floats"
	"testing"
)

func TestTrianglesIntersectionArea(t *testing.T) {
	a := Triangle{{0, 0}, {1, 0}, {0, 1}}
	b := Triangle{{0, 0}, {1, 1}, {0, 1}}
	e := 0.25
	if g := TrianglesIntersectionArea(a, b); f.Sign(g-e) != 0 {
		t.Errorf("Expected area %f, got %f.", e, g)
	}
}
