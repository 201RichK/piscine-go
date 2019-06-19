package piscine

func MakeRange(min, max int) []int {

	if min < max {
		tab := make([]int, max-min)
		j := 0
		for i := min; i < max; i++ {
			tab[j] = i
			j++
		}
		return tab
	} else {
		var tab []int = nil
		return tab
	}

}
