package piscine

func Rot14(str string) string {

	newStr := []byte(str)

	for i := 0; i < len(str); i++ {
		if (newStr[i] >= 'A' && newStr[i] <= 'L') || (newStr[i] >= 'a' && newStr[i] <= 'l') {
			newStr[i] += 14
		} else if (newStr[i] >= 'M' && newStr[i] <= 'Z') || (newStr[i] >= 'm' && newStr[i] <= 'z') {
			newStr[i] -= 12
		}
	}
	return string(newStr)
}
