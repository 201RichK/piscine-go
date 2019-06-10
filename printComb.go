package new

import "fmt"

func PrintComb() {
	a := 0
	b := 0
	c := 0
	for a <= 7 {
		for b <= 8 {
			for c <= 9 {

				if a < b && b < c {
					if a != 7 || b != 8 || c != 9 {
						fmt.Println(',')
						fmt.Println(' ')
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
