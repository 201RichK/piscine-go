package piscine

func IterativeFactorial(nb int) int {

	result := 1

	for index := 1; index < nb+1; index++ {
		result = result * index
	}

	return result

}
