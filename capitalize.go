package piscine

func Capitalize(s string) string {
	str := []byte(s)
	word := false
	for i, c := range str {
		if IsAlpha(c) {
			if !word {
				word = true
				str[i] = toUpper(c)
			} else {
				str[i] = toLower(c)
			}
		} else {
			word = false
		}
	}
	return string(str)
}
