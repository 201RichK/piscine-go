package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {

	} else if len(arguments) > 1 {
		filename := arguments[1]

		f, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(f))
		}
	}
}
