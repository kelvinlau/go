package ptree

func arrayCmp(a []int, x, y int) int {
	switch {
	case a[x] < a[y]:
		return -1
	case a[x] > a[y]:
		return +1
	default:
		return x - y
	}
}

type arrayPerm struct {
	a, d []int
}

func (p *arrayPerm) Len() int {
	return len(p.a)
}

func (p *arrayPerm) Swap(i, j int) {
	d := p.d
	d[i], d[j] = d[j], d[i]
}

func (p *arrayPerm) Less(i, j int) bool {
	a := p.a
	d := p.d
	x, y := d[i], d[j]
	return arrayCmp(a, x, y) < 0
}
