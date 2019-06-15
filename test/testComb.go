package main

import (
	"fmt"
)

func AppendRange(min, max int) []int {

	var tab []int

	if min < max {
		for i := min; i < max; i++ {
			tab = append(tab, i)
		}
	}

	return tab
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
