package geometry2d

import "fmt"

// Example_DynamicHalfPlane is the solution for
// http://codeforces.com/problemset/problem/379/E.
func Example_DynamicHalfPlane() {
	polylines := [][]float64{
		{7, 5, 5, 5, 9, 10, 9, 8, 7, 5, 10},
		{4, 2, 8, 2, 9, 1, 2, 8, 10, 7, 10},
		{9, 7, 7, 2, 5, 1, 5, 4, 7, 9, 7},
		{7, 3, 2, 10, 6, 9, 10, 2, 4, 2, 4},
		{1, 4, 8, 6, 9, 2, 1, 3, 6, 2, 8},
		{2, 4, 10, 7, 1, 1, 7, 9, 8, 9, 8},
		{8, 3, 10, 9, 4, 9, 9, 1, 9, 6, 3},
		{8, 10, 7, 2, 6, 2, 1, 3, 9, 7, 5},
		{1, 8, 3, 1, 7, 2, 8, 8, 3, 3, 9},
		{4, 1, 10, 2, 7, 4, 1, 9, 8, 1, 7}}
	k := len(polylines[0]) - 1
	area := make([]float64, len(polylines))
	for j := 0; j < k; j++ {
		p := NewDynamicPolygon(0, 0, 1, 2000)
		for i, l := range polylines {
			area1 := p.Area()
			hf := HalfPlane{Point{0, l[j]}, Point{1, l[j+1]}}
			p.Add(hf)
			area[i] += area1 - p.Area()
		}
	}
	for _, a := range area {
		fmt.Printf("%.12f\n", a)
	}
	// Output:
	// 71.500000000000
	// 6.500000000000
	// 4.133333333333
	// 3.419642857143
	// 0.636309523809
	// 2.267424242424
	// 0.646464646464
	// 2.375000000000
	// 0.000000000000
	// 0.000000000000
}
