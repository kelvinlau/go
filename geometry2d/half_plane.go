package geometry2d

import "sort"

// HalfPlane is the set of points on the positive side of a line.
type HalfPlane Line

// InHalfPlane returns true iff p is in a half plane, inclusive.
func InHalfPlane(hf HalfPlane, p Point) bool {
	return Sign(Cross(hf.P, hf.Q, p)) >= 0
}

// HalfPlaneIntersection returns the convex hull resulting from intersecting a
// set of half planes.
func HalfPlaneIntersection(hfs []HalfPlane) (ps []Point) {
	if len(hfs) < 2 {
		return
	}
	ls := []Line{}
	for _, hf := range hfs {
		ls = append(ls, Line(hf))
	}
	sort.Sort(LinesAngleComparator(ls))
	dq := []*Line{&ls[0], &ls[1]}
	for i := 2; i < len(ls); i++ {
		for len(dq) > 1 && !InHalfPlane(HalfPlane(ls[i]), IntersectionPoint(*dq[len(dq)-2], *dq[len(dq)-1])) {
			dq = dq[:len(dq)-1]
		}
		for len(dq) > 1 && !InHalfPlane(HalfPlane(ls[i]), IntersectionPoint(*dq[0], *dq[1])) {
			dq = dq[1:]
		}
		dq = append(dq, &ls[i])
	}
	for f := true; f; {
		f = false
		for len(dq) > 1 && !InHalfPlane(HalfPlane(*dq[0]), IntersectionPoint(*dq[len(dq)-2], *dq[len(dq)-1])) {
			dq = dq[:len(dq)-1]
			f = true
		}
		for len(dq) > 1 && !InHalfPlane(HalfPlane(*dq[len(dq)-1]), IntersectionPoint(*dq[0], *dq[1])) {
			dq = dq[1:]
			f = true
		}
	}
	for i := 0; i < len(dq); i++ {
		j := (i + 1) % len(dq)
		ps = append(ps, IntersectionPoint(*dq[i], *dq[j]))
	}
	return
}
