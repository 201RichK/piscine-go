package piscine

import "sort"

func Abort(a, b, c, d, e int) int {
	tbl := []int{a, b, c, d, e}
	sort.Ints(tbl)
	mediane := tbl[2]

	return mediane
}
