package main

import (
	"fmt"
)

func IsPrintable(str string) bool {
	b := []byte(str)
	var t bool
	for i := 0; i < len(b); i++ {
		if 65 <= str[i] && str[i] <= 90 || 97 <= str[i] && str[i] <= 122 || 47 <= str[i] && str[i] <= 57 {
			return true
		} else {
			t = false
		}
	}
	return t
}

func main() {
	fmt.Println(IsPrintable("01,02,03"))
	fmt.Println(IsPrintable("\f\t\f\f\t\n\v\a\t\v\t\r\b\r\a\t\t\f\t\t"))
}
