package my_test

import (
	. "."
	"testing"
)

func TestFuncA(t *testing.T) {
	expected := 123
	actual := FuncA()

	if expected != actual {
		t.Errorf("want %d, got %d", expected, actual)
	}
}

func BenchmarkFuncA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuncA()
	}
}
