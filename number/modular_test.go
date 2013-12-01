package number

import "testing"

func TestModularPower(t *testing.T) {
	if ModularPower(2, 4, 11) != 5 {
		t.Fatalf("2^4%11 should be %d, got %d.", 5, ModularPower(2, 4, 11))
	}
}

func TestModularInvert(t *testing.T) {
	if ModularInvert(2, 11) != 6 {
		t.Fatalf("2^-1%11 should be %d, got %d.", 6, ModularInvert(2, 11))
	}
}
