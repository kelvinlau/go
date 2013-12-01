package fft

import (
	"number"
	"testing"
)

func TestNewPoly(t *testing.T) {
	p := NewPoly([]int64{1, 0, 1})
	t.Logf("p: %#v", p)
	if len(p) != 8 {
		t.Fatalf("Wrong length: %d", len(p))
	}
}

func TestMul(t *testing.T) {
	p := NewPoly([]int64{1, 0, 1, 2})
	q := NewPoly([]int64{1, 1, 2, 3})
	r := Mul(p, q)
	e := NewPoly([]int64{1, 0, 2, 6})
	t.Logf("p: %#v", r)
	testEquals(r, e, t)
}

func TestGolden(t *testing.T) {
	if !number.IsPrime(modular) {
		t.Fatalf("%d is not a prime", modular)
	}
	if !number.IsPrimitiveRoot(modular, omega) {
		t.Fatalf("%d is not a primitive root of %d", omega, modular)
	}
}

func TestFFT(t *testing.T) {
	x1 := NewPoly(Poly{1, 1, 1, 1})
	x2 := NewPoly(Poly{1, 1, 1, 1})
	e3 := Poly{1, 2, 3, 4, 3, 2, 1, 0}
	y1 := FFT(x1)
	y2 := FFT(x2)
	t.Logf("y1: %#v", y1)
	t.Logf("y2: %#v", y2)
	y3 := Mul(y1, y2)
	t.Logf("y3: %#v", y3)
	x3 := IFFT(y3)
	t.Logf("x3: %#v", x3)
	testEquals(x3, e3, t)
}

func testEquals(a, b Poly, t *testing.T) {
	if len(a) != len(b) {
		t.Fatalf("Wrong length: %d", len(a))
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Fatalf("Wrong element %d: got %d, want %d", i, a[i], b[i])
		}
	}
}
