package piscine

func IsPrintable(str string) bool {
	b := []byte(str)
	var t bool
	for i := 0; i < len(b); i++ {
		if 65 <= str[i] && str[i] <= 90 || 97 <= str[i] && str[i] <= 122 || 47 <= str[i] && str[i] <= 57 || str[i] == 33 {
			return true
		} else {
			t = false
		}
	}
	return t
}
