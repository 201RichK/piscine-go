package piscine

func IsPrintable(str string) bool {

	for i := 0; i < len(str); i++ {
		if 0x20 <= str[i] && str[i] <= 0x7E {
			return false
		}
	}
	return true
}
