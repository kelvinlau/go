package geometry2d

import (
	. "github.com/kelvinlau/go/floats"
	"sort"
)

// HalfPlane is the set of points on the positive side of a line.
type HalfPlane Line

// HalfPlanes is a sort.Interface the sorts lines by angle.
type HalfPlanes []HalfPlane

func (hfs HalfPlanes) Len() int { return len(hfs) }
func (hfs HalfPlanes) Less(i, j int) bool {
	u, v := hfs[i], hfs[j]
	a1 := Angle(u.Q, u.P)
	a2 := Angle(v.Q, v.P)
	c := Sign(a1 - a2)
	return c < 0 || c == 0 && Sign(Cross(u.P, u.Q, v.P)) < 0
}
func (hfs HalfPlanes) Swap(i, j int) { hfs[i], hfs[j] = hfs[j], hfs[i] }

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
	sort.Sort(HalfPlanes(hfs))
	dq := []*HalfPlane{&hfs[0], &hfs[1]}
	ip := func(a, b *HalfPlane) Point {
		return IntersectionPoint(Line(*a), Line(*b))
	}
	for i := 2; i < len(hfs); i++ {
		for len(dq) > 1 && !InHalfPlane(hfs[i], ip(dq[len(dq)-2], dq[len(dq)-1])) {
			dq = dq[:len(dq)-1]
		}
		for len(dq) > 1 && !InHalfPlane(hfs[i], ip(dq[0], dq[1])) {
			dq = dq[1:]
		}
		dq = append(dq, &hfs[i])
	}
	for f := true; f; {
		f = false
		for len(dq) > 1 && !InHalfPlane(*dq[0], ip(dq[len(dq)-2], dq[len(dq)-1])) {
			dq = dq[:len(dq)-1]
			f = true
		}
		for len(dq) > 1 && !InHalfPlane(*dq[len(dq)-1], ip(dq[0], dq[1])) {
			dq = dq[1:]
			f = true
		}
	}
	for i := 0; i < len(dq); i++ {
		j := (i + 1) % len(dq)
		ps = append(ps, ip(dq[i], dq[j]))
	}
	return
}
