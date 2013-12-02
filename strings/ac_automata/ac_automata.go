package ac_automata

type node struct {
	next map[rune]*node
	fail *node
	word int
}

// An AcAutomata is a DFA that surpports multi-string matching.
type AcAutomata struct {
	root *node
}

// Build builds an AcAutomata from s slice of strings.
func Build(s []string) *AcAutomata {
	return nil
}

// Run reports how many string in the given set are substring of s.
func (a *AcAutomata) Run(s string) int {
	return 0
}
