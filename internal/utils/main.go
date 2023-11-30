package utils

import (
	"os"
	"strings"
)

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

func CountBytes(bytes []byte) int {
	var byteCount int

	for i := 0; i < len(bytes); i++ {
		byteCount++
	}

	return byteCount
}

func CountChars(bytes []byte) int {
	var charCount int

	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		arrLine := strings.Split(line, "")
		for i := 0; i < len(arrLine); i++ {
			charCount++
		}

		// New line counts as one cah so add it here
		charCount++
	}

	// Last line will not contain a newline so remove
	charCount--

	return charCount
}

func FileToBytes(fileName string) []byte {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return bytes
}

func stripTabAndSpaceFromLine(text string) []string {
	var flatWords []string

	words := strings.Split(text, " ")
	for _, word := range words {
		tabInWord := strings.Contains(word, "	")
		if tabInWord {
			splitWords := strings.Split(word, "	")

			for _, splitWord := range splitWords {
				flatWords = append(flatWords, splitWord)
			}
		} else {
			flatWords = append(flatWords, word)
		}
	}

	return flatWords
}

func CountWords(bytes []byte) int {
	var wordCount int

	fileContent := string(bytes)
	text := strings.Split(fileContent, "\n")

	for _, line := range text {
		line = strings.TrimSpace(line)

		if line != "" {
			words := stripTabAndSpaceFromLine(line)

			if words[0] != "" {
				wordCount += len(words)
			}
		}
	}

	return wordCount
}

