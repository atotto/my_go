package my_test

import (
	"testing"
)

func TestMyFuncA(t *testing.T) {

}

func BenchmarkMyFuncA(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
