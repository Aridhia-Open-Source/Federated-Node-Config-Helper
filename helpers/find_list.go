package helpers

func FindInList(item string, lst []string) int {
	for idx, val := range lst {
		if val == item {
			return idx
		}
	}
	return -1
}
