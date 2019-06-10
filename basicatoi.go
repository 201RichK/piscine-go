package piscine

func BasicAtoi(s string) int {

	i := 0
	for j := 0; j < len(s); j++ {

		i = i * 10
		i = i + int(s[j]) - '0'

	}
	return i
}
