package piscine

import "math"

func RecursiveFactorial(nb int) int {
	result := 1

	if nb >= 1 {
		result = nb * RecursiveFactorial(nb-1)

		if result > math.MaxInt32 {
			result = 0
		}

	} else if nb == 0 || nb == 1 {
		result = 1
	} else {
		result = 0
	}

	return result
}
