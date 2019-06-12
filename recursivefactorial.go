package piscine

import "math"

func RecursiveFactorial(nb int) int {
	if nb >= 1 {
		if result > math.MaxInt32 {
			return 0
		}
		return nb * RecursiveFactorial(nb-1)
	} else if nb == 0 || nb == 1 {
		return 1
	}
}
