package badinit

import (
	"fmt"
)

func init() {
	fmt.Println("bad init")
}

func main() {
	fmt.Println("hello")
}
