package treap

import "testing"

func TestTreap(t *testing.T) {
	treap := New()
	for x := 0; x < 100; x += 3 {
		treap.Insert(x)
	}
	for x := 95; x >= 0; x -= 5 {
		if treap.Find(x) == nil {
			treap.Insert(x)
		}
	}
	for x := 0; x < 100; x += 7 {
		treap.Erase(x)
	}

	e := []int{}
	for x := 0; x < 100; x++ {
		if (x%3 == 0 || x%5 == 0) && x%7 != 0 {
			e = append(e, x)
		}
	}

	if len(e) != treap.Size() {
		t.Fatalf("Expected size %v, %v.", len(e), treap.Size())
	}

	check := func(g []int) {
		if len(e) != len(g) {
			t.Fatalf("Expected %v, %v.", e, g)
		}
		for i := range e {
			if e[i] != g[i] {
				t.Fatalf("Expected %v, %v.", e, g)
			}
		}
	}

	g := []int{}
	treap.Each(func(x *Node) {
		g = append(g, x.Key)
	})
	check(g)

	g2 := []int{}
	for x := treap.Head(); x != nil; x = x.Next() {
		g2 = append(g2, x.Key)
	}
	check(g2)

	for x := -10; x <= 110; x++ {
		lb, rb := -1, -1
		for _, y := range e {
			if y >= x && lb == -1 {
				lb = y
			}
			if y > x && rb == -1 {
				rb = y
			}
		}
		glb := treap.LowerBound(x)
		grb := treap.UpperBound(x)

		check := func(fn string, x, e int, g *Node) {
			ok := false
			if e == -1 {
				ok = g == nil
			} else {
				ok = g != nil && g.Key == e
			}
			if !ok {
				t.Errorf("%s %d: expected %d, got %v.", fn, x, e, g)
			}
		}
		check("LowerBound", x, lb, glb)
		check("UpperBound", x, rb, grb)
	}

	for i, x := range e {
		if g := treap.Kth(i); g == nil && g.Key != x {
			t.Errorf("Kth %d: exptected %d, got %v.", i, e, g)
		}
	}
	for _, i := range []int{-1, len(e)} {
		if g := treap.Kth(i); g != nil {
			t.Errorf("Kth %d: exptected nil, got %v.", i, g)
		}
	}

	for i, x := range e {
		if g := treap.Count(x); g != i {
			t.Errorf("Count %d: exptected %d, got %v.", x, i, g)
		}
	}
	for _, x := range []int{-1, -100, -3, 0} {
		if g := treap.Count(x); g != 0 {
			t.Errorf("Count %d: exptected 0, got %v.", x, g)
		}
	}
	for _, x := range []int{100, 200, 3050} {
		if g := treap.Count(x); g != len(e) {
			t.Errorf("Count %d: exptected %d, got %v.", x, len(e), g)
		}
	}
}
