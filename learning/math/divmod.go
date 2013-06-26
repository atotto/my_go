package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(7)
	y := big.NewInt(3)

	q, r := new(big.Int).DivMod(x, y, new(big.Int))

	fmt.Printf("quotient:%d, remainder:%d\n", q, r)
}
