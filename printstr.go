package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	for i := 0; i <= len(str)-1; i++ {
		//fmt.Printf("%s", str[i])
		z01.PrintRune(rune(str[i]))
	}
}
