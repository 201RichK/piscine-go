package main

import (
	"fmt"
)

func IsPrintable(str string) bool {
	b := []byte(str)
	var t bool
	for i := 0; i < len(b); i++ {
		if 65 <= str[i] && str[i] <= 90 {
			return true
		} else {
			t = false
		}
	}
	return t
}

func main() {
	fmt.Println(IsPrintable(":)dKk`\"oipWtr"))
	fmt.Println(IsPrintable("\f\t\f\f\t\n\v\a\t\v\t\r\b\r\a\t\t\f\t\t"))
}
