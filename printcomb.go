package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {

	var a int = 0
	var b int = 0
	var c int = 0

	var sep string = ", "

	for a <= 7 {
		for b <= 8 {
			for c <= 9 {

				if a < b && b < c {

					if a == 7 && b == 8 && c == 9 {
						z01.PrintRune(rune(a))
						z01.PrintRune(rune(b))
						z01.PrintRune(rune(c))

					} else {
						z01.PrintRune(rune(a))
						z01.PrintRune(rune(b))
						z01.PrintRune(rune(c))
						z01.PrintRune(sep)
					}

				}

				c++
			}
			c = 0
			b++
		}

		b = 0
		a++
	}

}
