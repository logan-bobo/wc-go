package cmd

func CountLines(bytes []byte) int {
	var lineCount int

	for _, byte := range bytes {
		// The decimal representation of a UTF-8 newline is 10.
		if byte == 10 {
			lineCount++
		}
	}

	return lineCount
}
