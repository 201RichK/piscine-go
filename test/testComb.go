package main

import (
	"fmt"
)

func PrintComb() {

	var b int = 0
	var c int = 0

	//var sep string = ", "

	for b < 10 {
		for c < 10 {

			if b == c {

			} else {
				if b == 9 && c == 8 {
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

func main() {

	PrintComb()

}
