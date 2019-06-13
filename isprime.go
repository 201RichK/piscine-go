package piscine

func IsPrime(nb int) bool {
	var a bool = true
	if nb <= 1 {
		a = false
	} else {
		for i := 2; i <= nb-1; i++ {
			if nb%i == 0 {
				a = false
			}
		}
	}
	return a
}
