package main


import ( "github.com/01-edu/z01"
	"os"
)


func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr % 2 == 0 {
		return true 
	} else {
		return false 
	}
}

func main() {

	arguments := os.Args
	lengthOfArg := len(arguments) - 1

	OddMsg  := "I have an odd number of arguments"
	EvenMsg := "I have an even number of arguments"
	

	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}



