package piscine

func IterativeFactorial(int nb) int {

	result := 1

	for index := 1; index < nb+1; index++ {
		result = result * index
	}

	return result

}
