package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) == 4 {
		val1 := args[1]
		op := args[2]
		val2 := args[3]

		if _, err := strconv.Atoi(val1); err == nil {
			if _, err := strconv.Atoi(val2); err == nil {

				val1, _ := strconv.ParseInt(val1, 10, 64)
				val2, _ := strconv.ParseInt(val2, 10, 64)

				switch op {
				case "+":
					rslt := val1 + val2
					fmt.Println(rslt)
					break
				case "-":
					rslt := val1 - val2
					fmt.Println(rslt)
					break
				case "%":
					if val2 == 0 {
						fmt.Println("No Modulo by 0")
					} else {
						rslt := val1 % val2
						fmt.Println(rslt)
					}
					break
				case "*":
					rslt := val1 * val2
					fmt.Println(rslt)
					break
				case "/":
					if val2 == 0 {
						fmt.Println("No division by 0")
					} else {
						rslt := val1 / val2
						fmt.Println(rslt)
					}
					break
				default:
					fmt.Println("0")
				}
			}

		}
	}

}
