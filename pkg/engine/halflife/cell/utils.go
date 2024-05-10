package cell

func addToRight(left, right []string) []string {
	var lines []string
	for i := range len(left) {
		lines = append(lines, left[i]+right[i])
	}
	return lines
}

func addToDown(top, down []string) []string {
	return append(top, down...)
}
