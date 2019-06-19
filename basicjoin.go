package piscine

func BasicJoin(strs []string) string {

	var newWord string

	for i := 0; i < len(strs); i++ {
		newWord += strs[i]
	}
	return newWord
}
