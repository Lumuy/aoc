package intcode_program

func SliceContains(slice []string, ele string) bool {
	for _, s := range slice {
		if s == ele {
			return true
		}
	}

	return false
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
