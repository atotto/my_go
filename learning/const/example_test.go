package main

import (
	"fmt"
)

const (
	Const1 int = 1 << iota
	Const2
	Const3
	Const4
)

// go test -run Example
func ExampleConst() {
	fmt.Println(Const1)
	fmt.Println(Const2)
	fmt.Println(Const3)
	fmt.Println(Const4)
	// Output:
	// 1
	// 2
	// 4
	// 8
}
