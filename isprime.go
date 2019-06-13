package piscine

func IsPrime(nb int) bool {
	//a := 2 + (2 * (RecursiveFactorial(nb - 1)) % nb)
	var a bool
	if nb <= 1 {
		a = false
	}
	for i := 2; i <= int(nb); i++ {
		if nb%i == 0 {
			a = false
		} else {
			a = true
		}
	}
	return a
}
