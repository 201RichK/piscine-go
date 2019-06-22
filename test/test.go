package main

import (
	"fmt"
	"sort"
)

func Abort(a, b, c, d, e int) int {
	tbl := []int{a, b, c, d, e}
	newTbl := sort.IntSlice(tbl)
	mediane := newTbl[3]

	return mediane
}

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
