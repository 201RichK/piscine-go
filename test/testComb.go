package main

import (
	"fmt"
	"strings"
)

func Capitalize(s string) string {
	var newStr string
	for i := 0; i < len(s); i++ {
		if s[i] == " " {
			i = 0
		} else {
			newStr = strings.ToUpper(s[i])
		}
	}
}

func main() {
	fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
}
