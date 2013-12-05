package geometry2d

import "testing"

func testEqual(t *testing.T, es, qs []Point) {
	ok := true
	if len(qs) != len(es) {
		ok = false
	}
	for i := 0; ok && i < len(qs); i++ {
		if qs[i].X != es[i].X || qs[i].Y != es[i].Y {
			ok = false
		}
	}
	if !ok {
		t.Fatalf("Expected %v, got %v.", es, qs)
	}
}

func TestAngularSort(t *testing.T) {
	ps := []Point{
		{2, 2},
		{0, 0},
		{1, 1},
		{0, 2},
		{2, 0},
	}
	es := []Point{
		{0, 0},
		{2, 0},
		{1, 1},
		{2, 2},
		{0, 2},
	}
	AngularSort(ps)
	testEqual(t, es, ps)
}

func TestConvexHull(t *testing.T) {
	ps := []Point{
		{2, 2},
		{0, 0},
		{1, 1},
		{0, 2},
		{2, 0},
	}
	es := []Point{
		{0, 0},
		{2, 0},
		{2, 2},
		{0, 2},
	}
	qs := ConvexHull(ps)
	testEqual(t, es, qs)
}
