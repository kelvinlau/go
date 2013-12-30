// Package lctree implements link-cut tree.
package lctree

// Node is a node in the link-cut tree.
type Node struct {
	l, r, p, q *Node
	z          bool
	f, g, w    int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func up(x *Node) {
	l, r := x.l, x.r
	x.g = x.w
	if l != nil {
		x.g = max(x.g, l.g+l.f)
	}
	if r != nil {
		x.g = max(x.g, r.g+r.f)
	}
}

func down(x *Node) *Node {
	if x.z {
		x.z = false
		x.l, x.r = x.r, x.l
		if x.l != nil {
			x.l.z = !x.l.z
		}
		if x.r != nil {
			x.r.z = !x.r.z
		}
	}
	if x.f > 0 {
		if x.l != nil {
			x.l.f += x.f
		}
		if x.r != nil {
			x.r.f += x.f
		}
		x.w += x.f
		x.g += x.f
		x.f = 0
	}
	return x
}

func push(x *Node) {
	if x.p != nil {
		push(x.p)
	}
	down(x)
}

func lr(x *Node) {
	y := x.r
	b := y.l
	x.r = b
	if b != nil {
		b.p = x
	}
	y.p = x.p
	if x.p != nil {
		if x.p.l == x {
			x.p.l = y
		} else {
			x.p.r = y
		}
	}
	y.l = x
	x.p = y
	y.q = x.q
	x.q = nil
	up(x)
}

func rr(x *Node) {
	y := x.l
	b := y.r
	x.l = b
	if b != nil {
		b.p = x
	}
	y.p = x.p
	if x.p != nil {
		if x.p.l == x {
			x.p.l = y
		} else {
			x.p.r = y
		}
	}
	y.r = x
	x.p = y
	y.q = x.q
	x.q = nil
	up(x)
}

func rot(x *Node) {
	if x.p.l == x {
		rr(x.p)
	} else {
		lr(x.p)
	}
}

func splay(x *Node) {
	push(x)
	for x.p != nil {
		p := x.p
		if p.p == nil {
			rot(x)
		} else if (p.l == x) != (p.p.l == p) {
			rot(x)
			rot(x)
		} else {
			rot(p)
			rot(x)
		}
	}
	up(x)
}

func access(x *Node) {
	var u, v *Node = x, nil
	for u != nil {
		splay(u)
		if u.r != nil {
			u.r.q = u
			u.r.p = nil
		}
		u.r = v
		if v != nil {
			v.p = u
			v.q = nil
		}
		up(u)
		v = u
		u = u.q
	}
	splay(x)
}

// Rotate makes x the root of the tree.
func Rotate(x *Node) {
	access(x)
	x.z = !x.z
}

// Cut cuts out subtree of x from its parent.
func Cut(x *Node) {
	access(x)
	if x.l != nil {
		x.l.p = nil
		x.l = nil
		up(x)
	}
}

// Link add an edge from u to v. v must be a root.
func Link(u, v *Node) {
	access(u)
	access(v)
	v.l = u
	u.p = v
	up(v)
}

// Root returns the root of x.
func Root(x *Node) *Node {
	access(x)
	for down(x).l != nil {
		x = x.l
	}
	return x
}

// Parent returns the parent of x.
func Parent(x *Node) *Node {
	access(x)
	x = x.l
	if x == nil {
		return nil
	}
	for down(x).r != nil {
		x = x.r
	}
	return x
}

// Add let weight += d for nodes on path from x to y.
func Add(x, y *Node, d int) {
	Rotate(x)
	access(y)
	y.f += d
}

// Query returns max weight of nodes on path from x to y.
func Query(x, y *Node) int {
	Rotate(x)
	access(y)
	return y.g
}
