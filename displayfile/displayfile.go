package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args

	if len(arguments) > 2 {
		fmt.Println("Too many arguments")
	} else if len(arguments) == 1 {
		fmt.Println("File name missing")
	} else {
		fileName := arguments[1]
		f, _ := os.Open(fileName)

		arr := make([]byte, 14)

		f.Read(arr)
		fmt.Println(string(arr))
		f.Close()
	}
}
