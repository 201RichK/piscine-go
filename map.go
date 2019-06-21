package piscine

func Map(f func(int) bool, arr []int) []bool {
	var tbl make([]bool, leftLength)
	i := 0
	for _, elmt := range arr {
		tbl[i] = f(elmt)
		i++
	}
	return tbl
}
