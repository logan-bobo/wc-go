package wcutils

func CountLines(bytes []byte) int {
	var lineCount int

	for _, byteDecimal := range bytes {
		// The decimal representation of a UTF-8 newline is 10.
		if byteDecimal == 10 {
			lineCount++
		}
	}

	return lineCount
}
