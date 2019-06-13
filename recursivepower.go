package piscine

func RecursivePower(nb int, power int) int {
	result := 1
	if power < 0 {
		result = 0
	} else if power > 0 {

		result = nb * RecursivePower(nb, (power-1))
	} else {
		result = 1
	}

	return result
}
