package treap

import (
	"fmt"
	"math/rand"
)

// Treap is a balanced binary tree.
type Treap struct {
	root *Node
}

// New returns an empty treap.
func New() *Treap {
	return &Treap{
		root: nil,
	}
}

// Node is a node in the treap.
type Node struct {
	l, r, p *Node
	size, t int
	Key     int
}

func size(x *Node) int {
	if x != nil {
		return x.size
	}
	return 0
}

// Size reports the number of nodes in the treap.
func (t *Treap) Size() int {
	return size(t.root)
}

func rotate(x *Node) {
	y := x.p
	g := y.p
	x.p = g
	if g != nil {
		if g.l == y {
			g.l = x
		} else {
			g.r = x
		}
	}
	y.p = x
	if x == y.l {
		y.l = x.r
		if x.r != nil {
			x.r.p = y
		}
		x.r = y
	} else {
		y.r = x.l
		if x.l != nil {
			x.l.p = y
		}
		x.l = y
	}
	y.size = size(y.l) + size(y.r) + 1
	x.size = size(x.l) + size(x.r) + 1
}

func adjust(x *Node) *Node {
	for x.p != nil && x.t < x.p.t {
		rotate(x)
	}
	return up(x)
}

func insert(p, x *Node) *Node {
	if p == nil {
		return x
	}
	if x.Key < p.Key {
		p.l = insert(p.l, x)
		p.l.p = p
	} else {
		p.r = insert(p.r, x)
		p.r.p = p
	}
	p.size++
	return p
}

func up(x *Node) *Node {
	if x == nil {
		return nil
	}
	for x.p != nil {
		x = x.p
	}
	return x
}

// Insert inserts a record into the treap.
func (t *Treap) Insert(key int) {
	x := &Node{
		Key:  key,
		size: 1,
		t:    rand.Int(),
	}
	t.insertNode(x)
}

func (t *Treap) insertNode(x *Node) {
	insert(t.root, x)
	t.root = adjust(x)
}

// Find returns the node from a given key, or nil if not found.
func (t *Treap) Find(key int) *Node {
	p := t.root
	for p != nil {
		if p.Key == key {
			return p
		}
		if key < p.Key {
			p = p.l
		} else {
			p = p.r
		}
	}
	return nil
}

// Erase erases a node with the given key, return true on succ.
func (t *Treap) Erase(key int) bool {
	x := t.Find(key)
	if x == nil {
		return false
	}
	for x.l != nil || x.r != nil {
		if x.l != nil && (x.r == nil || x.l.t < x.r.t) {
			rotate(x.l)
		} else {
			rotate(x.r)
		}
	}
	y := x.p
	x.p = nil
	if y != nil {
		if y.l == x {
			y.l = nil
		} else {
			y.r = nil
		}
		for z := y; z != nil; z = z.p {
			z.size--
		}
	}
	t.root = up(y)
	return true
}

// Count return the number of nodes with key less than the given key.
func (t *Treap) Count(key int) int {
	x := t.root
	c := 0
	for x != nil {
		if x.Key >= key {
			x = x.l
		} else {
			c += size(x.l) + 1
			x = x.r
		}
	}
	return c
}

// Kth returns the kth node in the treap, or nil if out of bound.
func (t *Treap) Kth(k int) *Node {
	x := t.root
	for x != nil {
		if k == size(x.l) {
			return x
		} else if k < size(x.l) {
			x = x.l
		} else {
			k -= size(x.l) + 1
			x = x.r
		}
	}
	return nil
}

func search(x *Node, key int, eq bool) *Node {
	if x == nil {
		return nil
	}
	if key < x.Key || key == x.Key && eq {
		if y := search(x.l, key, eq); y != nil {
			return y
		}
		return x
	}
	return search(x.r, key, eq)
}

// LowerBound return the smallest node that >= given key.
func (t *Treap) LowerBound(key int) *Node {
	return search(t.root, key, true)
}

// UpperBound return the smallest node that > given key.
func (t *Treap) UpperBound(key int) *Node {
	return search(t.root, key, false)
}

// Head returns the smallest node.
func (t *Treap) Head() *Node {
	x := t.root
	if x == nil {
		return nil
	}
	for x.l != nil {
		x = x.l
	}
	return x
}

// Next returns the next node of a given node in-order.
func (x *Node) Next() *Node {
	if x.r != nil {
		x = x.r
		for x.l != nil {
			x = x.l
		}
		return x
	} else {
		for x.p != nil && x.p.r == x {
			x = x.p
		}
		return x.p
	}
}

// Each feeds all nodes in-order to a given function.
func (t *Treap) Each(f func(x *Node)) {
	var dfs func(x *Node)
	dfs = func(x *Node) {
		if x == nil {
			return
		}
		dfs(x.l)
		f(x)
		dfs(x.r)
	}
	dfs(t.root)
}

// Print prints the data in the tree in-order.
func (t *Treap) Print() {
	t.Each(func(x *Node) {
		fmt.Printf("%v, ", x.Key)
	})
	fmt.Println()
}
