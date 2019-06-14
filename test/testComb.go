package main

import (
	"fmt"
)

func IsAlpha(str string) bool {

	var word bool

	for i := 0; i < len(str); i++ {
		if 'a' <= str[i] && str[i] <= 'z' || 'A' <= str[i] && str[i] <= 'Z' || '0' <= str[i] && str[i] <= '9' {
			word = true
		} else {
			return false
		}
	}
	return word
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}
