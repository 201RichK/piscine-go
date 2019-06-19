package piscine

func IsNumeric(str string) bool {

	var word bool

	for i := 0; i < len(str); i++ {
		if '0' <= str[i] && str[i] <= '9' {
			word = true
		} else {
			return false
		}
	}
	return word
}
