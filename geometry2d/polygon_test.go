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

func TestRotateCalipers(t *testing.T) {
	ps := []Point{
		{0, 0},
		{1, 0},
		{2, 1},
		{0, 1},
	}
	area1, peri1 := 2.0, 6.0
	area, peri := RotateCalipers(ps)
	if Sign(area-area1) != 0 {
		t.Errorf("Expected area %f, got %f.", area1, area)
	}
	if Sign(peri-peri1) != 0 {
		t.Errorf("Expected peri %f, got %f.", peri1, peri)
	}
}

func TestTriangulate(t *testing.T) {
	ps := []Point{
		{0, 0},
		{2, 0},
		{2, 1},
		{1, 1},
		{1, 2},
		{0, 2},
	}
	ts := Triangulate(ps)
	if len(ts) != len(ps)-2 {
		t.Errorf("Expected #triangles %d, got %d.", len(ps)-2, len(ts))
	}
	area := 0.0
	for _, t := range ts {
		area += AreaPolygon(t[:])
	}
	e := AreaPolygon(ps)
	if area != e {
		t.Errorf("Expected sum of triangle areas %f, got %f.", e, area)
	}
}

func TestPolygonIntersectionArea(t *testing.T) {
	ps := []Point{
		{1, 0},
		{2, 0},
		{2, 3},
		{1, 3},
	}
	qs := []Point{
		{0, 1},
		{3, 1},
		{3, 2},
		{0, 2},
	}
	e := 1.0
	if g := PolygonIntersectionArea(ps, qs); Sign(g-e) != 0 {
		t.Errorf("Expected area %f, got %f.", e, g)
	}
}
