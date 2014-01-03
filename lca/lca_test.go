package lca

import (
	"testing"
)

func TestLca(t *testing.T) {
	tree := NewTree(10)
	tree.Link(3, 9)
	tree.Link(9, 5)
	tree.Link(1, 8)
	tree.Link(8, 5)
	tree.Link(7, 4)
	tree.Link(6, 0)
	tree.Link(3, 6)
	tree.Link(6, 2)
	tree.Link(0, 7)
	tree.Build()
	testLca := func(u, v, e int) {
		if g := tree.Lca(u, v); g != e {
			t.Errorf("Lca(%d, %d): expected %d, got %d.", u, v, e, g)
		}
	}
	testDist := func(u, v, e int) {
		if g := tree.Dist(u, v); g != e {
			t.Errorf("Dist(%d, %d): expected %d, got %d.", u, v, e, g)
		}
	}
	testLca(4, 2, 0)
	testLca(9, 9, 9)
	testLca(1, 0, 0)
	testLca(3, 2, 6)
	testDist(2, 3, 2)
	testDist(1, 1, 0)
	testDist(4, 1, 8)
	testDist(5, 9, 1)
	testDist(2, 4, 4)
	testDist(9, 7, 4)
}
