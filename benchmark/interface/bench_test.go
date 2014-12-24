package example_test

import "testing"

var testdata_int64 []int64
var testdata_interface []interface{}
var N = 10000

func init() {
	testdata_int64 = make([]int64, N)
	testdata_interface = make([]interface{}, N)
	for i := 0; i < N; i++ {
		val := int64(i)
		testdata_int64[i] = val
		testdata_interface[i] = val
	}
}

func BenchmarkInt(b *testing.B) {
	recvdata := make([]int64, N)
	for i := 0; i < b.N; i++ {
		for n, val := range testdata_int64 {
			recvdata[n] = val
		}
	}
}

func BenchmarkInterface(b *testing.B) {
	recvdata := make([]interface{}, N)
	for i := 0; i < b.N; i++ {
		for n, val := range testdata_interface {
			recvdata[n] = val
		}
	}
}
