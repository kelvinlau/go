// Package blossom implements Edmonds Blossom Contraction Algorithm.
package blossom

// MaxMatches calculates the maximum matches of a graph given the adj matrix.
// Returns m: the max num of matches;
//   match[i] == j mean i and j match, or -1 if not match for i.
func MaxMatches(g [][]bool) (m int, match []int) {
	n := len(g)
	match = make([]int, n)
	base := make([]int, n)
	father := make([]int, n)
	inblossom := make([]bool, n)
	inque := make([]bool, n)
	queue := make([]int, n)
	q := queue[:0]

	push := func(u int) {
		inque[u] = true
		q = append(q, u)
	}

	pop := func() int {
		u := q[0]
		q = q[1:]
		return u
	}

	lca := func(u, v int) int {
		vt := make([]bool, n)
		for u != -1 {
			u = base[u]
			vt[u] = true
			if match[u] == -1 {
				break
			}
			u = father[match[u]]
		}
		for v != -1 {
			v = base[v]
			if vt[v] {
				return v
			}
			v = father[match[v]]
		}
		return -1
	}

	reset := func(u, anc int) {
		for u != anc {
			v := match[u]
			inblossom[base[v]] = true
			inblossom[base[u]] = true
			v = father[v]
			if base[v] != anc {
				father[v] = match[u]
			}
			u = v
		}
	}

	contract := func(u, v int) {
		for i := 0; i < n; i++ {
			inblossom[i] = false
		}
		anc := lca(u, v)
		reset(u, anc)
		reset(v, anc)
		if base[u] != anc {
			father[u] = v
		}
		if base[v] != anc {
			father[v] = u
		}
		for i := 0; i < n; i++ {
			if inblossom[base[i]] {
				base[i] = anc
				if !inque[i] {
					push(i)
				}
			}
		}
	}

	find := func(start int) int {
		for i := 0; i < n; i++ {
			father[i] = -1
			inque[i] = false
			base[i] = i
		}
		q = queue[:0]
		push(start)
		for len(q) > 0 {
			u := pop()
			for v := 0; v < n; v++ {
				if g[u][v] && base[v] != base[u] && match[v] != u {
					if v == start || (match[v] != -1 && father[match[v]] != -1) {
						contract(u, v)
					} else if father[v] == -1 {
						father[v] = u
						if match[v] != -1 {
							push(match[v])
						} else {
							return v
						}
					}
				}
			}
		}
		return -1
	}

	augment := func(u int) {
		for u != -1 {
			v := father[u]
			w := match[v]
			match[u] = v
			match[v] = u
			u = w
		}
	}

	m = 0
	for i := range match {
		match[i] = -1
	}
	for i := 0; i < n; i++ {
		if match[i] == -1 {
			if end := find(i); end != -1 {
				augment(end)
				m++
			}
		}
	}
	return
}
