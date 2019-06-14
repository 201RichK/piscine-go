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
	fmt.Println(IsAlpha("P8rB#0~XDj%ss"))
	fmt.Println(IsAlpha(":ukSJSbnIxCGO"))
	fmt.Println(IsAlpha("3_rI?xN/zNoDn"))
	fmt.Println(IsAlpha("je"))
}
