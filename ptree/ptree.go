// Package ptree implements persistent tree that supports range lower bound query.
package ptree

import "sort"

const (
	Inf int = 0x7fffffff
)

// PTree is a persistent binary search tree that supports range lower bound
// queries.
type PTree struct {
	a    []int
	n    int
	root []*node
}

// New constructs a PTree from a slice of ints.
func New(a []int) *PTree {
	n := len(a)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = i
	}
	sort.Sort(&arrayPerm{a, b})
	t := &PTree{}
	t.a = a
	t.n = len(a)
	t.root = make([]*node, n+1)
	t.root[n] = t.build(b)
	for i := n - 1; i >= 0; i-- {
		t.root[i] = t.del(t.root[i+1], i)
	}
	return t
}

// RangeLB returns the minimum element in the range [l, r) that >= v,
// or returns Inf if no element found.
func (t *PTree) RangeLB(l, r, v int) int {
	if l < 0 {
		l = 0
	}
	if r > t.n {
		r = t.n
	}
	if l >= r {
		return Inf
	}
	if x := t.lb(t.root[l], t.root[r], v); x != nil {
		return t.a[x.d]
	}
	return Inf
}

type node struct {
	d, s int
	l, r *node
}

func leaf(x *node) bool {
	return x.l == nil && x.r == nil
}

func (t *PTree) build(b []int) *node {
	if len(b) == 1 {
		return &node{
			d: b[0],
			s: 1,
		}
	}
	m := len(b) / 2
	return &node{
		d: b[m],
		s: len(b),
		l: t.build(b[:m]),
		r: t.build(b[m:]),
	}
}

func (t *PTree) del(x *node, d int) *node {
	if leaf(x) {
		return nil
	}
	if arrayCmp(t.a, d, x.d) < 0 {
		l := t.del(x.l, d)
		if l == nil {
			return x.r
		}
		return &node{
			l: l,
			r: x.r,
			d: x.d,
			s: x.s - 1,
		}
	} else {
		r := t.del(x.r, d)
		if r == nil {
			return x.l
		}
		return &node{
			l: x.l,
			r: r,
			d: x.d,
			s: x.s - 1,
		}
	}
}

func (t *PTree) lb(x, y *node, v int) *node {
	if leaf(y) {
		if t.a[y.d] >= v && x == nil {
			return y
		}
		return nil
	}
	var xl, xr *node
	switch {
	case x == nil:
		xl, xr = nil, nil
	case arrayCmp(t.a, x.d, y.d) == 0:
		xl, xr = x.l, x.r
	case arrayCmp(t.a, x.d, y.d) < 0:
		xl, xr = x, nil
	case arrayCmp(t.a, x.d, y.d) > 0:
		xl, xr = nil, x
	}
	if t.count(y.l, v) > t.count(xl, v) {
		return t.lb(xl, y.l, v)
	} else {
		return t.lb(xr, y.r, v)
	}
}

func (t *PTree) count(x *node, v int) int {
	if x == nil {
		return 0
	}
	if leaf(x) {
		if t.a[x.d] >= v {
			return 2
		}
		return 0
	}
	if v <= t.a[x.d] {
		return t.count(x.l, v) + x.r.s
	}
	return t.count(x.r, v)
}
