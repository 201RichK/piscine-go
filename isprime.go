package piscine

import "math"

func IsPrime(nb int) bool {
	a := 2 + (2 * (RecursiveFactorial(nb - 1)) % nb)

	switch a {
	case 2:
		return false
	default:
		return true
	}
}
func RecursiveFactorial(nb int) int {
	if nb > 1 {
		if nb > math.MaxInt32 {
			return 0
		}
		return nb * RecursiveFactorial(nb-1)
	} else if nb == 0 || nb == 1 {
		return 1
	}

	return 0
}
