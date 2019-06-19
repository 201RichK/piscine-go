package piscine

func Join(strs []string, sep string) string {
	var newWord string

	for i := 0; i < len(strs); i++ {
		if i == (len(strs) - 1) {
			newWord += strs[i]
		} else {
			newWord += strs[i] + sep
		}
	}
	return newWord
}
