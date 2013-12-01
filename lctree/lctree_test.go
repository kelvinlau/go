package lctree

import (
	"math/rand"
	"testing"
)

func TestLctree(t *testing.T) {
	//   0      4
	//   |      |
	//   1      5
	//  / \
	// 2   3
	nodes := [6]Node{}
	a := func(i int) *Node {
		return &nodes[i]
	}
	testEqual := func(a, b *Node) {
		if a != b {
			t.Fatalf("Expected %p, got %p.", a, b)
		}
	}
	Link(a(0), a(1))
	Link(a(1), a(2))
	Link(a(1), a(3))
	Link(a(4), a(5))
	testEqual(Root(a(0)), a(0))
	testEqual(Root(a(1)), a(0))
	testEqual(Root(a(2)), a(0))
	testEqual(Root(a(3)), a(0))
	testEqual(Root(a(4)), a(4))
	testEqual(Root(a(5)), a(4))
	testEqual(Parent(a(0)), nil)
	testEqual(Parent(a(1)), a(0))
	testEqual(Parent(a(2)), a(1))
	testEqual(Parent(a(3)), a(1))
	testEqual(Parent(a(4)), nil)
	testEqual(Parent(a(5)), a(4))

	Cut(a(1))
	Cut(a(5))
	Rotate(a(2))
	Link(a(0), a(5))
	Link(a(0), a(4))
	Link(a(0), a(2))
	//   0
	//  /|\
	// 4 5 2
	//     |
	//     1
	//     |
	//     3
	testEqual(Root(a(0)), a(0))
	testEqual(Root(a(1)), a(0))
	testEqual(Root(a(2)), a(0))
	testEqual(Root(a(3)), a(0))
	testEqual(Root(a(4)), a(0))
	testEqual(Root(a(5)), a(0))
	testEqual(Parent(a(0)), nil)
	testEqual(Parent(a(1)), a(2))
	testEqual(Parent(a(2)), a(0))
	testEqual(Parent(a(3)), a(1))
	testEqual(Parent(a(4)), a(0))
	testEqual(Parent(a(5)), a(0))

	testQuery := func(i, j, e int) {
		if v := Query(a(i), a(j)); v != e {
			t.Fatalf("Query(%d, %d) should be %d, got %d.", i, j, e, v)
		}
	}
	Add(a(4), a(2), 1)
	Add(a(5), a(1), 2)
	Add(a(3), a(2), 4)
	testQuery(4, 5, 3)
	testQuery(4, 2, 7)
	testQuery(3, 3, 4)
	testQuery(0, 0, 3)
}

func BenchmarkLctree(b *testing.B) {
	n := 10000
	nodes := make([]Node, n)
	a := func(i int) *Node {
		return &nodes[i]
	}
	for k := 0; k < b.N; k++ {
		x := a(rand.Intn(n))
		y := a(rand.Intn(n))
		if Root(x) == Root(y) {
			Cut(x)
		} else {
			Rotate(y)
			Link(x, y)
		}
	}
}
