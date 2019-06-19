package piscine

func ConcatParams(args []string) string {

	rtr := ""

	for i, res := range args {
		rtr += string(res)
		if i != len(args)-1 {
			rtr += "\n"
		}
	}

	return rtr
}
