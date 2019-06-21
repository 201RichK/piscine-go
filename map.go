package piscine 

func Map(f func(int) bool, arr []int) []bool {
	tbl := []bool
	i := 0
	for _, elmt := range arr {
		tbl[i] = f(elmt)
		i++
	}
}