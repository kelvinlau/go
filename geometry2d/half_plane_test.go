package geometry2d

import "testing"

func TestHalfPlaneIntersection(t *testing.T) {
	a := Point{0, 0}
	b := Point{1, 0}
	c := Point{1, 1}
	d := Point{0, 1}
	hfs := []HalfPlane{
		{a, b},
		{c, d},
		{b, c},
		{d, a},
		{a, c},
	}
	ps := HalfPlaneIntersection(hfs)
	area := AreaPolygon(ps)
	e := 0.5
	if area != e {
		t.Errorf("Expected area %f, got %f.", e, area)
	}
}
