package piscine

func Rot14(str string) string {

	newStr := []byte(str)

	for i := 0; i < len(str); i++ {
		if newStr[i] == 'M' || newStr[i] == 'N' {
			newStr[i] -= 12
		}
		if (newStr[i] >= 'A' && newStr[i] < 'L') || (newStr[i] >= 'a' && newStr[i] < 'l') {
			newStr[i] += 14
		} else if (str[i] >= 'O' && str[i] <= 'Z') || (str[i] >= 'o' && str[i] <= 'z') {
			newStr[i] -= 14
		}
	}
	return string(newStr)
}
