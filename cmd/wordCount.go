package cmd

import (
	"strings"
)

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
