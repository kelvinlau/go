package lca

import (
	"unsafe"
)

type Tree struct {
	n  int
	lg int
	ns []node
}

type node struct {
	dep int
	pnt []*node
	adj []*node
}

func New(n int) *Tree {
	lg := 1
	for d := 1; d < n; d <<= 1 {
		lg++
	}
	t := &Tree{
		n:  n,
		lg: lg,
		ns: make([]node, n),
	}
	for i := range t.ns {
		x := &t.ns[i]
		x.dep = -1
		x.pnt = make([]*node, lg)
	}
	return t
}

func (t *Tree) Link(u, v int) {
	x := &t.ns[u]
	y := &t.ns[v]
	x.adj = append(x.adj, y)
	y.adj = append(y.adj, x)
}

func (t *Tree) Build() {
	t.build(&t.ns[0], &t.ns[0], 0)
	for d := 1; d < t.lg; d++ {
		for i := range t.ns {
			x := &t.ns[i]
			x.pnt[d] = x.pnt[d-1].pnt[d-1]
		}
	}
}

func (t *Tree) build(u, p *node, d int) {
	u.pnt[0] = p
	u.dep = d
	for _, v := range u.adj {
		if v.dep == -1 {
			t.build(v, u, d+1)
		}
	}
}

func (t *Tree) Lca(u, v int) int {
	x := &t.ns[u]
	y := &t.ns[v]
	z := t.lca(x, y)
	return int((uintptr(unsafe.Pointer(z)) - uintptr(unsafe.Pointer(&t.ns[0]))) / unsafe.Sizeof(t.ns[0]))
}

func (t *Tree) lca(x, y *node) *node {
	if x.dep < y.dep {
		x, y = y, x
	}
	d := x.dep - y.dep
	for i := 0; i < t.lg; i++ {
		if d&(1<<uint(i)) > 0 {
			x = x.pnt[i]
		}
	}
	for i := t.lg - 1; i >= 0; i-- {
		x1 := x.pnt[i]
		y1 := y.pnt[i]
		if x1 != y1 {
			x, y = x1, y1
		}
	}
	if x != y {
		x = x.pnt[0]
	}
	return x
}

func (t *Tree) Dist(u, v int) int {
	x := &t.ns[u]
	y := &t.ns[v]
	z := t.lca(x, y)
	return x.dep + y.dep - 2*z.dep
}
