package piscine

func ConcatParams(args []string) string {

	rtr := ""

	for i := 0; i < len(args); i++ {
		rtr += args[i] + "\n"
	}

	return rtr
}
