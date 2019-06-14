package piscine

func IsAlpha(str string) bool {

	var word bool

	for i := 0; i < len(str); i++ {
		if 'a' <= str[i] && str[i] <= 'z' || 'A' <= str[i] && str[i] <= 'Z' || '0' <= str[i] && str[i] <= '9' {
			word = true
		} else {
			word = false
		}
	}
	return word
}
