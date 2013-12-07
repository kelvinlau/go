package geometry2d

// InscribedCircleCenter returns the center of the max circle that fits into
// the triangle abc.
func InscribedCircleCenter(A, B, C Point) Point {
	a := Dist(B, C)
	b := Dist(C, A)
	c := Dist(A, B)
	p := a + b + c
	return Point{(a*A.X + b*B.X + c*C.X) / p, (a*A.Y + b*B.Y + c*C.Y) / p}
}

// CircumCircleCenter returns the center of the min circle that contains the
// triangle abc.
func CircumCircleCenter(A, B, C Point) Point {
	a1 := B.X - A.X
	b1 := B.Y - A.Y
	c1 := (Sqr(a1) + Sqr(b1)) / 2
	a2 := C.X - A.X
	b2 := C.Y - A.Y
	c2 := (Sqr(a2) + Sqr(b2)) / 2
	d := a1*b2 - a2*b1
	return Point{A.X + (c1*b2-c2*b1)/d, A.Y + (a1*c2-a2*c1)/d}
}
