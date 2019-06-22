package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	a := 1
	for i := 1; i < len(args); i++ {
		if args[i] == "01" || args[i] == "galaxy" || args[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		} else {
			a++
			if len(args) == a {
				fmt.Println("")
			}
		}
	}
}
