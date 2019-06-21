package piscine

func Any(f func(string) bool, arr []string) bool {

	var type bool = false

	for i := 0; i < len(arr); i++ {
		if f(i) {
			type = true
		}
	}

	return type
}