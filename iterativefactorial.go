package piscine

import "math"

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 0 {
		for i := 1; i < nb+1; i++ {
			result = result * i
			if result > math.MaxInt32 {
				return 0
			}
		}
		return result
	}
}
