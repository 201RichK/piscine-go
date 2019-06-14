package piscine

func isPrintable(c byte) bool {
	return 0x20 <= c && c <= 0x7E
}

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if !isPrintable(s[i]) {
			return false
		}
	}
	return true
}
