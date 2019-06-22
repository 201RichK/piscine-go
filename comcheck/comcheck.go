package piscine

import "fmt"

func Comcheck() {
	args := os.args

	taille := len(args)

	for i := 1; i < len(args); i++ {
		if args[i] == "01" || args[i] == "galaxy" {
			return fmt.Println("Alert!!!")
		} else {
			fmt.Println("")
		}
	}
}
