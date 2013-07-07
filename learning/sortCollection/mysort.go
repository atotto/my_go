package mysort

import ()

/* We should implement sort.Interface:
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less returns whether the element with index i should sort
    // before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
*/

//
type StringLengthSort []string

// Len is part of sort.Interface.
func (s StringLengthSort) Len() int {
	return len(s)
}

// Swap is part of sort.Interface.
func (s StringLengthSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is part of sort.Interface.
func (s StringLengthSort) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
