package my_test

import (
	. "."
	"fmt"
)

func ExampleFuncA_exampleA() {
	m := FuncA()
	fmt.Printf("result: %d", m)
	// Output: result: 123
}

func ExampleFuncA_exampleB() {
	m := FuncA()
	fmt.Println(m)
	// Output:
	// 123
}
