package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	} else if s[0] == '-' {
		return BasicAtoi2(s[1:]) * -1
	} else if s[0] == '+' {
		return BasicAtoi2(s[1:])
	} else {
		return BasicAtoi2(s)
	}
}
