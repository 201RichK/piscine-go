package main

import (
	"fmt"
)

func IterativePower(nb int, power int) int {

	if power < 0 {
		return 0
	} else if power >= 0 {

		result := 1

		for i := 0; i < power; i++ {
			result *= nb
		}
		return result
	}
	return 0
}

func main() {
	arg1 := 4
	arg2 := 0
	fmt.Println(IterativePower(arg1, arg2))
}
