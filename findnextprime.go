package piscine

func FindNextPrime(nb int) int {

	prime := nb
	for i := 0; i <= 10; i++ {
		if IsPrime(prime) {
			return prime
			break
		}
		prime++
	}
	return 0
}
