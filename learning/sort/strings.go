package main

import( 
	"fmt"
	"sort"
)

func main() {
	s := []string {"ccc", "ddd", "bbb", "aaa"}

	sort.Strings(s)
	fmt.Println(s)
}