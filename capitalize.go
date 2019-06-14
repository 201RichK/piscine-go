package piscine

import "strings"

func Capitalize(s string) string {
	str := strings.Title(ToLower(s))
	idx := Index(s, "_")
	if idx == -1 {
		return str
	} else {
		str = strings.Replace(str, string(str[idx+1]), ToUpper(string(str[idx+1])), idx+1)
		return str
	}
}
