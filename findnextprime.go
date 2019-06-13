package piscine

func FindNextPrime(nb int) int {
	a := nb

	for i := nb + 1; i > nb; i++ {
		if IsPrime(a) {
			return a
		}
		a++
	}
	return a
}
