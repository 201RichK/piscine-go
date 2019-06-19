package piscine

/*
	rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, to distinguish character values from integer values


	byte is an alias for uint8 and is equivalent to uint8 in all ways. It is used, by convention, to distinguish byte values from 8-bit unsigned integer values.

	https://golang.org/src/bytes/example_test.goz
*/

func BasicAtoi2(s string) int {
	for _, value := range []byte(s) {
		if value < '0' || value > '9' {
			return 0
		}
	}
	return BasicAtoi(s)
}
