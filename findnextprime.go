package piscine

func FindNextPrime(nb int) int {

	a := nb
	for i := nb; i >= nb; i++ {
		if IsPrime(i) {
			a = i
			return i
		}

	}
	return a
}
