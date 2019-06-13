package piscine

func FindNextPrime(nb int) int {
	a := nb - 1
	i := nb + 1
	for i > nb {
		a++
		if IsPrime(a) {
			return a
		}
		i++
	}
	return a
}
