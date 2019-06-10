package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintComb() {

	var b int = 0
	var c int = 0

	//var sep string = ", "

	for b <= 8 {
		for c <= 9 {

			if b < c {

				if a == 7 && b == 8 && c == 9 {
					fmt.Printf("%d%d", b, c)
				} else {
					fmt.Printf("%d%d", b, c)
					//fmt.Println(a, b, c)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}

			}

			c++
		}
		c = 0
		b++
	}

	z01.PrintRune('\n')

}
