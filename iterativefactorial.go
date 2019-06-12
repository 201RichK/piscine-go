package piscine

func IterativeFactorial(nb int) int {

	result := 1

	i := 1
	for i < nb+1 {
		result = result * i

		i++
	}
	return result

}
