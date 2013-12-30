// Package segtree implements a kind of Segment Tree.
// It supports 2 operations:
//   Inc(l, r, y): increment a[l, r) by y;
//   Max():				 returns the max(a[0, n)).
package segtree

type SegTree struct {
	n       int
	f, g, x []int
}

// NewSegTree make a new segment tree with n zeros.
func NewSegTree(n int) *SegTree {
	t := &SegTree{
		n: n,
		f: make([]int, 2*n),
		g: make([]int, 2*n),
		x: make([]int, 2*n),
	}
	t.build(0, n)
	return t
}

func (t *SegTree) build(u, v int) {
	i := id(u, v)
	t.x[i] = u
	if u+1 < v {
		d := mid(u, v)
		t.build(u, d)
		t.build(d, v)
	}
}

func id(u, v int) int {
	t := u + v - 1
	if u+1 < v {
		t |= 1
	}
	return t
}

func mid(u, v int) int {
	return (u + v + 1) >> 1
}

// Inc increments [a, b) by y.
func (t *SegTree) Inc(a, b, y int) {
	n := t.n
	f := t.f
	g := t.g
	x := t.x
	var inc func(u, v int)
	inc = func(u, v int) {
		i := id(u, v)
		if b <= u || v <= a {
			return
		}
		if a <= u && v <= b {
			f[i] += y
			g[i] += y
			return
		}
		d := mid(u, v)
		l := id(u, d)
		r := id(d, v)
		inc(u, d)
		inc(d, v)
		if f[l] > f[r] {
			f[i] = g[i] + f[l]
			x[i] = x[l]
		} else {
			f[i] = g[i] + f[r]
			x[i] = x[r]
		}
	}
	inc(0, n)
}

// Max returns the max element position and value.
func (t *SegTree) Max() (x, y int) {
	i := id(0, t.n)
	x, y = t.x[i], t.f[i]
	return
}
