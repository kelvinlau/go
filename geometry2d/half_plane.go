package geometry2d

import (
	. "github.com/kelvinlau/go/floats"
	"github.com/kelvinlau/go/treap"
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

// InHalfPlaneStrict returns true iff p is in a half plane, exclusive.
func InHalfPlaneStrict(hf HalfPlane, p Point) bool {
	return Sign(Cross(hf.P, hf.Q, p)) > 0
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

// DynamicPolygon maintains a polygon that accepts dynamic half plane
// intersection operations.
type DynamicPolygon struct {
	t     *treap.Treap
	area2 float64
}

// NewDynamicPolygon returns a new DynamicPolygon with the given surrounding
// rectangle.
func NewDynamicPolygon(x1, y1, x2, y2 float64) *DynamicPolygon {
	t := treap.New(func(a, b interface{}) bool {
		u := a.(HalfPlane)
		v := b.(HalfPlane)
		a1 := Angle(u.P, u.Q)
		a2 := Angle(v.P, v.Q)
		return Sign(a1-a2) < 0
	})
	a := Point{x1, y1}
	b := Point{x2, y1}
	c := Point{x2, y2}
	d := Point{x1, y2}
	t.Insert(HalfPlane{a, b}, b)
	t.Insert(HalfPlane{b, c}, c)
	t.Insert(HalfPlane{c, d}, d)
	t.Insert(HalfPlane{d, a}, a)
	p := &DynamicPolygon{
		t:     t,
		area2: 2 * (x2 - x1) * (y2 - y1),
	}
	return p
}

// Add adjusts the dynamic polygon to the intersection of current state and hf.
func (p *DynamicPolygon) Add(hf HalfPlane) {
	t := p.t
	if t.Size() == 0 {
		return
	}

	prev := func(x *treap.Node) *treap.Node {
		y := x.Prev()
		if y == nil {
			y = t.Tail()
		}
		return y
	}
	next := func(x *treap.Node) *treap.Node {
		y := x.Next()
		if y == nil {
			y = t.Head()
		}
		return y
	}
	point := func(x *treap.Node) Point {
		return x.Val.(Point)
	}
	line := func(x *treap.Node) Line {
		return Line(x.Key.(HalfPlane))
	}

	b := t.LowerBound(hf)
	if b == nil {
		b = t.Head()
	}
	a := prev(b)
	pv := point(a)
	if InHalfPlane(hf, pv) {
		return
	}
	a, b = prev(a), a
	for t.Size() > 1 && !InHalfPlaneStrict(hf, point(a)) {
		p.area2 -= Cross2(point(a), point(b))
		t.Erase(b)
		a, b = prev(a), a
	}
	p.area2 -= Cross2(point(a), point(b))
	a = next(b)
	b = next(a)
	for t.Size() > 1 && !InHalfPlaneStrict(hf, point(a)) {
		p.area2 -= Cross2(pv, point(a))
		pv = point(a)
		t.Erase(a)
		a, b = b, next(b)
	}
	p.area2 -= Cross2(pv, point(a))
	if t.Size() > 1 {
		a, b = prev(a), a
		q := IntersectionPoint(Line(hf), line(b))
		t.Insert(hf, q)
		a.Val = IntersectionPoint(line(a), Line(hf))
		p.area2 += Cross2(q, point(b))
		p.area2 += Cross2(point(a), q)
		a, b = prev(a), a
		p.area2 += Cross2(point(a), point(b))
	} else {
		t.Clear()
		p.area2 = 0
	}
}

// Area returns the current area of the dynamic polygon.
func (p *DynamicPolygon) Area() float64 {
	return p.area2 / 2
}
