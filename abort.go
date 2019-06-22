package piscine

import "sort"

func Abort(a, b, c, d, e int) int {
	tbl := []int{a, b, c, d, e}
	newTbl := sort.IntSlice(tbl)
	mediane := newTbl[3]

	return mediane
}
