// Package ac implements Aho–Corasick multi-pattern matching algorithm.
package ac

type node struct {
	next map[rune]*node
	fail *node
	word int
}

func newNode() *node {
	return &node{next: make(map[rune]*node)}
}

// An Automata is a DFA that surpports multi-string matching.
type Automata struct {
	root *node
}

// New constructs an empty Automata.
func New() *Automata {
	return &Automata{root: newNode()}
}

// Insert inserts a string into the automata.
func (a *Automata) Insert(s string) {
	x := a.root
	for _, z := range s {
		if _, ok := x.next[z]; !ok {
			x.next[z] = newNode()
		}
		x = x.next[z]
	}
	x.word++
}

// Build finalize the automata after inserting strings.
func (a *Automata) Build() {
	q := []*node{a.root}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for z, v := range u.next {
			x := u.fail
			for x != nil && x.next[z] == nil {
				x = x.fail
			}
			if x != nil {
				v.fail = x.next[z]
			} else {
				v.fail = a.root
			}
			q = append(q, v)
		}
	}
}

// Run reports how many string in the given set are substring of s, can be run
// only once.
func (a *Automata) Run(s string) int {
	ans := 0
	x := a.root
	for _, z := range s {
		for x != nil && x.next[z] == nil {
			x = x.fail
		}
		if x != nil {
			x = x.next[z]
			for y := x; y != nil && y.word != -1; y = y.fail {
				ans += y.word
				y.word = -1
			}
		} else {
			x = a.root
		}
	}
	return ans
}
