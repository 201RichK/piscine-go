package main

import (
	"fmt"
)

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

func main() {
	fmt.Println(IsPrintable("UbbuGXt|KRdS1"))
	fmt.Println(IsPrintable("@vZ1nu$eoHfZJ"))

	//fmt.Println(string(0x20))
}
