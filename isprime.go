package piscine

import "math"

func IsPrime(nb int) bool {
	//a := 2 + (2 * (RecursiveFactorial(nb - 1)) % nb)
	var a bool
	for i := 2; i <= int(math.Floor(float64(nb)/2)); i++ {
		if nb%i == 0 {
			a = true
		} else {
			a = false
		}
	}
	return a
}
