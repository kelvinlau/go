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

func TestGCD(t *testing.T) {
	test := func(x, y, e int64) {
		if g := GCD(x, y); g != e {
			t.Fatalf("GCD of %d %d should be %d, got %d.", x, y, e, g)
		}
	}
	test(100, 26, 2)
	test(1, 20, 1)
	test(3, 4, 1)
	test(15, 25, 5)
}

func TestExGCD(t *testing.T) {
	test := func(x, y int64) {
		a, b, g := ExGCD(x, y)
		if a*x+b*y != g || g != GCD(x, y) {
			t.Fatalf("ExGCD of %d %d got %d, %d, %d, calculated gcd = %d.",
				x, y, a, b, g, a*x+b*y)
		}
	}
	test(100, 26)
	test(1, 20)
	test(3, 4)
	test(15, 25)
}
