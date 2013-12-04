package number

import "testing"

func TestModularPower(t *testing.T) {
	test := func(a, b, m, e int64) {
		if g := ModularPower(a, b, m); g != e {
			t.Fatalf("%d ^ %d %% %d should be %d, got %d.", a, b, m, e, g)
		}
	}

	test(2, 4, 11, 5)
	test(123456789123456, 321654987654321, 456789123456765, 243638565486786)
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
	test(3, 4, 1)
	test(1, 20, 1)
	test(15, 25, 5)
	test(100, 26, 2)
}

func TestExGCD(t *testing.T) {
	test := func(x, y int64) {
		a, b, g := ExGCD(x, y)
		if a*x+b*y != g || g != GCD(x, y) {
			t.Fatalf("ExGCD of %d %d got %d, %d, %d, calculated gcd = %d.",
				x, y, a, b, g, a*x+b*y)
		}
	}
	test(3, 4)
	test(1, 20)
	test(15, 25)
	test(100, 26)
}

func TestModularSystem(t *testing.T) {
	if x := ModularSystem([]int64{3, 5, 7}, []int64{1, 2, 3}); x != 52 {
		t.Fatalf("Expected %d, got %d.", 52, x)
	}
}

func TestModularLog(t *testing.T) {
	test := func(a, b, m int64) {
		if !IsPrime(m) {
			t.Fatalf("x%d", m)
		}
		if x := ModularLog(a, b, m); x != -1 {
			if c := ModularPower(a, x, m); c != b {
				t.Fatalf("ModularLog(%d, %d, %d) got %d, but %d^%d%%%d = %d.",
					a, b, m, x, a, x, m, c)
			}
		} else {
			t.Logf("ModularLog(%d, %d, %d) = -1.", a, b, m)
		}
	}
	test(3, 2, 7)
	test(13, 4, 17)
	test(23, 23, 100003)
	test(56, 7, 10007)
}
