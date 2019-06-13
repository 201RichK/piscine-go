package piscine

func IsPrime(nb int) bool {
	//a := 2 + (2 * (RecursiveFactorial(nb - 1)) % nb)
	var a bool = true
	if nb <= 1 {
		a = false
	} else {
		for i := 2; i <= nb; i++ {
			if nb%i == 0 {
				a = false
			}
		}
	}

	return a
}
