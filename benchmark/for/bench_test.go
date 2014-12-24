package example_test

import (
	"reflect"
	"testing"
)

var testdata []int64
var N = 10000

func init() {
	testdata = make([]int64, N)
	for i := 0; i < N; i++ {
		testdata[i] = int64(i)
	}
}

func BenchmarkFor(b *testing.B) {
	recvdata := make([]int64, N)
	for i := 0; i < b.N; i++ {
		for n := 0; n < N; n++ {
			recvdata[n] = testdata[n]
		}
	}
}

func BenchmarkForRange(b *testing.B) {
	recvdata := make([]int64, N)
	for i := 0; i < b.N; i++ {
		for n, val := range testdata {
			recvdata[n] = val
		}
	}
}

func BenchmarkForReflect(b *testing.B) {
	recvdata := make([]int64, N)
	v := reflect.ValueOf(testdata)
	length := v.Len()
	for i := 0; i < b.N; i++ {
		for n := 0; n < length; n++ {
			recvdata[n] = v.Index(n).Int()
		}
	}
}
