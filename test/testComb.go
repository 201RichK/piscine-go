package main

import (
	"fmt"
	"sort"
)

/*https://golang.org/pkg/sort/*/
func SortIntegerTable(table []int) {
	//	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(table)
	//fmt.Println(s)
}

func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}
