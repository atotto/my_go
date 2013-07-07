package mysort_test

import (
	. "."
	"fmt"
	"sort"
)

func ExampleMySort() {
	languages := []string{"Go", "Java", "JavaScript", "PHP", "Python"}

	sort.Sort(StringLengthSort(languages))
	fmt.Println(languages)

	// Output:
	// [Go PHP Java Python JavaScript]
}
