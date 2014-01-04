package blossom_test

import (
	"fmt"
	"github.com/kelvinlau/go/blossom"
)

type edge struct {
	u, v int
}

type testcase struct {
	n  int
	es []edge
}

// ExampleSeagrp is the solution for
// http://www.codechef.com/JAN14/problems/SEAGRP/
func ExampleSeagrp() {
	ts := []testcase{
		{2, []edge{{0, 1}, {0, 1}}},
		{3, []edge{{0, 1}, {1, 2}}},
		{4, []edge{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}, {2, 3}}},
	}

	for _, t := range ts {
		n := t.n
		g := make([][]bool, n)
		for i := 0; i < n; i++ {
			g[i] = make([]bool, n)
		}
		for _, e := range t.es {
			g[e.u][e.v] = true
			g[e.v][e.u] = true
		}

		k, _ := blossom.MaxMatches(g)
		if k*2 == n {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
	// Output:
	// YES
	// NO
	// YES
}
