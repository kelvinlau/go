package number

import "testing"

func TestPrimitiveRoot(t *testing.T) {
	for n := int64(2); n < 100; n++ {
		t.Logf("Primitive root of %d is %d.", n, PrimitiveRoot(n))
	}
	if PrimitiveRoot(6) != 0 {
		t.Fatalf("Primitive root of %d is %d, got %d.", 6, 0, PrimitiveRoot(6))
	}
	if PrimitiveRoot(90) != 11 {
		t.Fatalf("Primitive root of %d is %d, got %d.", 90, 11, PrimitiveRoot(90))
	}
}

func BenchmarkPrimitiveRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := int64(2); n < 10000; n++ {
			_ = PrimitiveRoot(n)
		}
	}
}
