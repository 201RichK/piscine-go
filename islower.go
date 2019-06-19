package piscine

func IsLower(str string) bool {

	var lower bool
	for i := 0; i < len(str); i++ {
		if 'a' <= str[i] && str[i] <= 'z' {
			lower = true
		} else {
			return false
		}
	}
	return lower
}
