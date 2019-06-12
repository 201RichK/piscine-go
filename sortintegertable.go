package piscine

import "sort"

/*https://golang.org/pkg/sort/*/
func SortIntegerTable(table []int) {
	//	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(table)
	//fmt.Println(s)
}
