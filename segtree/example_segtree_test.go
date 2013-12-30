// Solution for http://codeforces.com/contest/377/problem/D.
package segtree_test

import (
	"fmt"
	"sort"

	"github.com/kelvinlau/go/segtree"
)

type Worker struct {
	i, l, v, r int
}

type Event struct {
	t int
	*Worker
}

type Events []Event

func (a Events) Less(i, j int) bool {
	x, y := a[i], a[j]
	vx := x.v
	if x.t == -1 {
		vx = x.r
	}
	vy := y.v
	if y.t == -1 {
		vy = y.r
	}
	if vx != vy {
		return vx < vy
	}
	return x.t > y.t
}
func (a Events) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Events) Len() int      { return len(a) }

func Uniq(xs []int) []int {
	sort.Ints(xs)
	xx := xs[:0]
	for _, x := range xs {
		if len(xx) == 0 || xx[len(xx)-1] != x {
			xx = append(xx, x)
		}
	}
	return xx
}

// Example_segtree is the solution for
// http://codeforces.com/contest/377/problem/D.
func Example_segtree() {
	n := 4
	a := []Worker{
		{0, 2, 8, 9},
		{1, 1, 4, 7},
		{2, 3, 6, 8},
		{3, 5, 8, 10},
	}

	xs := make([]int, 2*n)
	for i := range a {
		xs[i+i+0] = a[i].l
		xs[i+i+1] = a[i].v
	}
	xx := Uniq(xs)

	xid := func(x int) int {
		return sort.Search(len(xx), func(i int) bool {
			return xx[i] >= x
		})
	}

	es := make([]Event, 2*n)
	for i := range a {
		es[i+i+0] = Event{+1, &a[i]}
		es[i+i+1] = Event{-1, &a[i]}
	}
	sort.Sort(Events(es))

	ymax, x1, x2 := 0, 0, 0
	t := segtree.NewSegTree(len(xx))
	for _, e := range es {
		l := xid(e.l)
		r := xid(e.v) + 1
		t.Inc(l, r, e.t)
		if x, y := t.Max(); y > ymax {
			ymax = y
			x1 = xx[x]
			x2 = e.v
		}
	}

	ls := []int{}
	for i := range a {
		if a[i].l <= x1 && x1 <= a[i].v && a[i].v <= x2 && x2 <= a[i].r {
			ls = append(ls, a[i].i+1)
		}
	}
	fmt.Println(ymax)
	pretty := fmt.Sprint(ls)
	fmt.Println(pretty[1 : len(pretty)-1])
	// Output:
	// 3
	// 1 3 4
}
