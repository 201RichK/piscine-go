package main

import (
	"fmt"
)

func iseven(n int) (b bool) {
	if n%2 == 0 {
		b = true
	} else {
		b = false
	}
	return
}

var steps int = 1

func CollatzCountdown(start int) int {

	if iseven(start) {
		steps += 1
		start = start / 2
		CollatzCountdown(start)
	} else {
		steps += 1
		start = start*3 + 1
		CollatzCountdown(start)
	}
	return steps
}

func main() {
	steps := CollatzCountdown(12)
	fmt.Println(steps)
}
