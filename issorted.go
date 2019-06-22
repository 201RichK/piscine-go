package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	a := true
	for i := 0; i < len(tab); i++ {
		if f(tab[i], tab[i+1]) >= 0 {
			a = true
		} else if f(tab[i], tab[i+1]) < 0 {
			return a
		}
	}
	return a
}
