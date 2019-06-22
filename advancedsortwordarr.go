package piscine

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for x := 0; x < len(array)-1; x++ {
		for y := x + 1; y < len(array); y++ {
			if f(array[x], array[y]) == 1 {
				arrTmp := array[x]
				array[x] = array[y]
				array[y] = arrTmp
			}
		}
	}
}
