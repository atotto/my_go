package main

import( 
	"fmt"
	"sort"
)

func main() {
	m := map[string]string{
		"bbb" : "2",
		"ccc" : "3",
		"aaa" : "1",
	}
	
    keys := make([]string, len(m))
    i := 0
    for k, _ := range m {
        keys[i] = k
        i++
    }
    sort.Strings(keys)

    fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}