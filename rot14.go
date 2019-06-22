package piscine

func Rot14(str string) string {

	newStr := []byte(str)

	for i := 0; i < len(str); i++ {
		if (newStr[i] >= 'A' && newStr[i] <= 'L') || (newStr[i] >= 'a' && newStr[i] <= 'l') {
			newStr[i] += 14
		} else if (str[i] >= 'M' && str[i] <= 'Z') || (str[i] >= 'o' && str[i] <= 'z') {
			newStr[i] -= 12
		}
	}
	return string(newStr)
}
