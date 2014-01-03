package geometry3d

import (
	. "github.com/kelvinlau/go/floats"
	"testing"
)

func TestNorm(t *testing.T) {
	v := Vector{3, -4, 5}
	u := Norm(v)
	if !Eq(Len(u), 1) {
		t.Errorf("Norm %v, got %v, len: %f.", v, u, Len(u))
	}
}

func TestRandPerp(t *testing.T) {
	v := Vector{3, -4, 5}
	u := RandPerp(v)
	if Sign(Dot(u, v)) != 0 {
		t.Errorf("RandPerp %v, got %v, dot product: %f.", v, u, Dot(u, v))
	}
}
