package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	arguments := os.Args

	sort.Strings(arguments)

	for i := 0; i < len(arguments); i++ {
		fmt.Println(arguments[i])
	}
}
