// Package fft implements fast fourier transform using number theory.
package fft

import "github.com/kelvinlau/go/number"

const (
	modular = 1004535809
	omega   = 3
)

// Next power of 2 >= n.
func nextPowerOf2(n int64) int64 {
	p := int64(1)
	for p < n {
		p <<= 1
	}
	return p
}

// Poly represents the coeffecients of a polynomial.
type Poly []int64

// NewPoly pad zeros to p, to have length of a power of 2 and >= 2*len(p).
func NewPoly(p Poly) Poly {
	l := int64(len(p))
	n := 2 * nextPowerOf2(l)
	for ; l < n; l++ {
		p = append(p, 0)
	}
	return p
}

// Mul calculates the product of p and q.
func Mul(p, q Poly) Poly {
	r := make(Poly, len(p))
	for i := range p {
		r[i] = p[i] * q[i] % modular
	}
	return r
}

func fft(v, y Poly, r, ws, s int64) {
	if len(y) == 1 {
		y[0] = v[0]
	} else {
		m := len(y) / 2
		fft(v[:], y[:m], r, ws*ws%r, s+s)
		fft(v[s:], y[m:], r, ws*ws%r, s+s)
		wsk := int64(1)
		for k1, k2 := 0, m; k1 < m; k1, k2 = k1+1, k2+1 {
			y[k1], y[k2] = y[k1]+wsk*y[k2], y[k1]-wsk*y[k2]
			y[k1] = (y[k1]%r + r) % r
			y[k2] = (y[k2]%r + r) % r
			wsk = wsk * ws % r
		}
	}
}

// FFT returns the fourier transform of v.
func FFT(v Poly) Poly {
	y := make(Poly, len(v))
	n := int64(len(v))
	w := number.ModularPower(omega, (modular-1)/n, modular)
	fft(v, y, modular, w, 1)
	return y
}

// IFFT returns the inverted fourier transform of v.
func IFFT(v Poly) Poly {
	y := make(Poly, len(v))
	n := int64(len(v))
	w := number.ModularPower(omega, (modular-1)/n, modular)
	w1 := number.ModularInvert(w, modular)
	fft(v, y, modular, w1, 1)
	n1 := number.ModularInvert(n, modular)
	for i := range y {
		y[i] = y[i] * n1 % modular
	}
	return y
}
