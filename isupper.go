package piscine

func IsUpper(str string) bool {

	var uper bool
	for i := 0; i < len(str); i++ {
		if 'A' <= str[i] && str[i] <= 'A' {
			uper = true
		} else {
			return false
		}
	}
	return uper
}
