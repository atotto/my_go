package divmod_test

import (
	"math/big"
	"testing"
)

func TestBigPackageDivMod(t *testing.T) {
	x := big.NewInt(7)
	y := big.NewInt(3)

	expected_q := big.NewInt(2)
	expected_r := big.NewInt(1)

	q, r := divmod(x, y)
	if q.Cmp(expected_q) != 0 || r.Cmp(expected_r) != 0 {
		t.Errorf("quotient=%d remainder=%d, but q=%d, r=%d\n", expected_q, expected_r, q, r)
	}
}

func TestMyDivMod(t *testing.T) {
	x := 7
	y := 3

	expected_q := 2
	expected_r := 1

	q, r := myDivMod(x, y)
	if q != expected_q || r != expected_r {
		t.Errorf("quotient=%d remainder=%d, but q=%d, r=%d\n", expected_q, expected_r, q, r)
	}
}

func BenchmarkBigPackageDivMod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := big.NewInt(7)
		y := big.NewInt(3)

		divmod(x, y)
	}
}

func BenchmarkMyDivMod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := 7
		y := 3

		myDivMod(x, y)
	}
}

func divmod(x, y *big.Int) (q, r *big.Int) {
	q, r = new(big.Int).DivMod(x, y, new(big.Int))
	return
}

func myDivMod(x, y int) (q, r int) {
	q = x / y
	r = x % y
	return
}
