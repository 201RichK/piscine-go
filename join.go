package piscine

func Join(strs []string, sep string) string {
	var newWord string

	for i := 0; i < len(strs); i++ {
		newWord += strs[i] + sep
	}
	return newWord
}
