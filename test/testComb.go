package main

import (
	"fmt"
)

func ConcatParams(args []string) string {

	rtr := ""

	for i := 0; i < len(args); i++ {
		rtr += args[i] + "\n"
	}

	return rtr
}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
