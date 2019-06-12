package piscine

import "math"

func RecursiveFactorial(nb int) int {
	result := 0

	if nb >= 1 {

		if result > math.MaxInt32 {
			return result
		}

		result = nb * RecursiveFactorial(nb-1)

	} else if nb == 0 || nb == 1 {
		result = 1
	}

	return result
}
