package main

import (
	"fmt"
)

func RecursivePower(nb int, power int) int {
	result := 1
	if power < 0 {
		result = 0
	} else if power > 0 {

		result = nb * RecursivePower(nb, (power-1))
	} else {
		result = 1
	}

	return result
}

func main() {
	arg1 := 4
	arg2 := 0
	fmt.Println(RecursivePower(arg1, arg2))
}
