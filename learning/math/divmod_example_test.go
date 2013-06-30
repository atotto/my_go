package divmod_test

import (
	"fmt"
	"math/big"
)

func ExampleInt_DivMod() {
	x := big.NewInt(7)
	y := big.NewInt(3)

	q, r := new(big.Int).DivMod(x, y, new(big.Int))

	fmt.Printf("quotient:%d, remainder:%d\n", q, r)
	// Output:
	// quotient:2, remainder:1
}
