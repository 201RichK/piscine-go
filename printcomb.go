package piscine

import "github.com/01-edu/z01"

func show(a int, b int, c int) {
	if a < b && b < c {
		printDigit(a)
		printDigit(b)
		printDigit(c)
		if a != 7 || b != 8 || c != 9 {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}

func PrintComb() {
	a := 0
	b := 0
	c := 0
	for a <= 7 {
		for b <= 8 {
			for c <= 9 {
				show(a, b, c)
				c++
			}
			c = 0
			b++
		}
		b = 0
		a++
	}
	z01.PrintRune('\n')
}
