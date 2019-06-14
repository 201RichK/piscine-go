package piscine

func IsPrintable(str string) bool {
	var uper bool = true
	for i := 0; i < len(str); i++ {

		if 'A' <= str[i] && str[i] <= 'Z' {
			uper = true
		} else if 'a' <= str[i] && str[i] <= 'z' {
			return false
		}
	}
	return uper
}
