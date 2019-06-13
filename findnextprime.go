package piscine

func FindNextPrime(nb int) int {
	a := nb - 1

	for i := nb + 1; i > nb; i++ {
		a++
		if IsPrime(a) {
			return a
		}
	}
	return a
}
