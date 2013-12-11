package geometry3d

import (
	f "github.com/kelvinlau/go/floats"
	"testing"
)

func TestNorm(t *testing.T) {
	v := Vector{3, -4, 5}
	u := Norm(v)
	if f.Sign2(Len(u), 1) != 0 {
		t.Errorf("Norm %v, got %v, len: %f.", v, u, Len(u))
	}
}

func TestRandPerp(t *testing.T) {
	v := Vector{3, -4, 5}
	u := RandPerp(v)
	if f.Sign(Dot(u, v)) != 0 {
		t.Errorf("RandPerp %v, got %v, dot product: %f.", v, u, Dot(u, v))
	}
}
