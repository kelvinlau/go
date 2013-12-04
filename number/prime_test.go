package number

import "testing"

func TestIsPrime(t *testing.T) {
	es := []int64{2, 3, 5, 7, 11, 13, 17, 19}
	ps := []int64{}
	for i := int64(1); i < 20; i++ {
		if IsPrime(i) {
			ps = append(ps, i)
		}
	}
	t.Logf("ps: %#v", ps)
	testEquals(ps, es, t)
}

func testEquals(a, b []int64, t *testing.T) {
	if len(a) != len(b) {
		t.Fatalf("Wrong length: %d", len(a))
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Errorf("Wrong element %d: got %d, want %d", i, a[i], b[i])
		}
	}
}
