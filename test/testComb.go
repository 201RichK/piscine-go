package main

import "fmt"

func main() {

	s := "Hello World!"
	//s = StrRev(s)
	fmt.Println(s)
	a := len([]rune(s))
	fmt.Println(a)

	for i := a; i > 0; i-- {
		fmt.Print(s[i])
	}

}
