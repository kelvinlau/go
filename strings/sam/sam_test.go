package sam

import "testing"

func TestSAM(t *testing.T) {
	s := New()
	strs := []string{"golang", "google", "apple", "microsoft", "awesome"}
	for _, str := range strs {
		s.Add(str)
	}
	suffixes := []string{"lang", "gle", "le", "e", "soft", "some"}
	nonsuffixes := []string{"a", "z", "mountain", "eagle"}
	for _, str := range suffixes {
		if s.Run(str) == false {
			t.Errorf("%s should be accepted.", str)
		}
	}
	for _, str := range nonsuffixes {
		if s.Run(str) == true {
			t.Errorf("%s should not be accepted.", str)
		}
	}
}
