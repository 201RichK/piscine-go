package piscine

func Any(f func(string) bool, arr []string) bool {

	var fType bool = false

	for i := 0; i < len(arr); i++ {
		if f(i) {
			fType = true
		}
	}

	return fType
}
