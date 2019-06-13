package piscine

func IsPrime(nb int) bool {
	a := 2 + (2 * (RecursiveFactorial(nb - 1)) % nb)

	switch a {
	case 2:
		return false
	default:
		return true
	}
}
