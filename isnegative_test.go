package piscine

import "fmt"

func IsNegative(n int) {
	if n < 0 {
		fmt.Println('T')
	} else {
		fmt.Println('F')
	}
	fmt.Println('\n')
}
