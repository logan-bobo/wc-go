package cmd

func CountBytes(bytes []byte) int {
	var byteCount int

	for i := 0; i < len(bytes); i++ {
		byteCount++
	}

	return byteCount
}
