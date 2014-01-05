// Package sam implements a suffix automata, a DFA accepts all suffixes of some
// strings.
package sam

type node struct {
	l    int
	t    bool
	next map[rune]*node
	fail *node
}

func empty(l int) *node {
	return &node{
		l:    l,
		next: make(map[rune]*node),
	}
}

func clone(y *node, l int) *node {
	x := &node{
		l:    l,
		t:    y.t,
		next: make(map[rune]*node),
		fail: y.fail,
	}
	for c, z := range y.next {
		x.next[c] = z
	}
	return x
}

// SAM is a suffix automata, a DFA accepts all suffixes of some strings.
type SAM struct {
	start *node
}

// New contructs an empty SAM.
func New() *SAM {
	return &SAM{
		start: empty(0),
	}
}

func (s *SAM) extend(end *node, c rune) *node {
	x := empty(end.l + 1)
	var p *node
	for p = end; p != nil && p.next[c] == nil; p = p.fail {
		p.next[c] = x
	}
	if p != nil {
		q := p.next[c]
		if p.l+1 < q.l {
			q1 := clone(q, p.l+1)
			x.fail = q1
			q.fail = q1
			for ; p != nil && p.next[c] == q; p = p.fail {
				p.next[c] = q1
			}
		} else {
			x.fail = q
		}
	} else {
		x.fail = s.start
	}
	return x
}

// Add adds a string into the SAM.
func (s *SAM) Add(str string) {
	x := s.start
	for _, c := range str {
		if x.next[c] != nil && x.next[c].l == x.l+1 {
			x = x.next[c]
		} else {
			x = s.extend(x, c)
		}
	}
	for ; x != s.start; x = x.fail {
		x.t = true
	}
}

// Run judge whether suffix is a suffix of one of the strings in SAM.
func (s *SAM) Run(suffix string) bool {
	x := s.start
	for _, c := range suffix {
		x = x.next[c]
		if x == nil {
			return false
		}
	}
	return x.t
}
