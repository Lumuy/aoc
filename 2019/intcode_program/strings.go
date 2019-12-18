package intcode_program

func SliceContains(slice []string, ele string) bool {
	for _, s := range slice {
		if s == ele {
			return true
		}
	}

	return false
}
