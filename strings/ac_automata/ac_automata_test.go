package ac_automata

import "testing"

func TestAcAutomata(t *testing.T) {
	ss := []string{
		"abc",
		"abc",
		"abd",
		"xxx",
		"xxxyy",
		"xyzz",
	}
	f := "123abcdefxxxyzz"
	e := 4

	a := New()
	for _, s := range ss {
		a.Insert(s)
	}
	a.Build()
	if g := a.Run(f); g != e {
		t.Errorf("Expected %d matches, got %d.", e, g)
	}
}
