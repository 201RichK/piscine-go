package piscine

import "strings"

func SplitWhiteSpaces(str string) []string {
	rtr := strings.Fields(str)
	return rtr
}
