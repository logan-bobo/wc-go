package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func countLines(bytes []byte) int {
	var lineCount int

	for _, byteDecimal := range bytes {
		// The decimal representation of a UTF-8 newline is 10.
		if byteDecimal == 10 {
			lineCount++
		}
	}

	return lineCount
}

func countBytes(bytes []byte) int {
	byteCount := len(bytes)

	return byteCount
}

func countChars(bytes []byte) int {
	var charCount int

	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		arrLine := strings.Split(line, "")
		for i := 0; i < len(arrLine); i++ {
			charCount++
		}

		// New line counts so add it
		charCount++
	}

	// Last line will not contain a newline so remove
	charCount--

	return charCount
}

func fileToBytes(fileName string) []byte {
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

func countWords(bytes []byte) int {
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

func main() {
	var fileName string
	var fileBytes []byte
	var flag string

	legalFlags := [4]string{"-c", "-l", "-w", "-m"}

	file := os.Stdin
	fileInfo, err := file.Stat()

	if err != nil {
		panic(err)
	}

	size := fileInfo.Size()
	if size > 0 {
		fileBytes, _ = io.ReadAll(os.Stdin)
		if len(os.Args) > 1 {
			for _, legalFlag := range legalFlags {
				if os.Args[1] == legalFlag {
					flag = os.Args[1]
				}
			}
		}

	} else {
		if len(os.Args) < 3 {
			for _, legalFlag := range legalFlags {
				if len(os.Args) < 2 {
					fmt.Println("No File provided")
					os.Exit(1)
				}

				if os.Args[1] == legalFlag {
					fmt.Println("Flag set with no filename")
					os.Exit(1)
				}
			}
			fileName = os.Args[1]
			fileBytes = fileToBytes(fileName)
		} else {
			for _, legalFlag := range legalFlags {
				if os.Args[1] == legalFlag {
					flag = os.Args[1]
					fileName = os.Args[2]
					fileBytes = fileToBytes(fileName)
				}
			}
		}
	}

	if flag == "-c" {
		fileByteTotal := countBytes(fileBytes)
		fmt.Println(fileByteTotal, fileName)

	} else if flag == "-l" {
		fileLines := countLines(fileBytes)
		fmt.Println(fileLines, fileName)

	} else if flag == "-w" {
		fileWords := countWords(fileBytes)
		fmt.Println(fileWords, fileName)

	} else if flag == "-m" {
		fileChars := countChars(fileBytes)
		fmt.Println(fileChars, fileName)

	} else {
		fileByteCount := countBytes(fileBytes)
		fileLines := countLines(fileBytes)
		fileWords := countWords(fileBytes)
		fmt.Println(fileLines, fileWords, fileByteCount, fileName)
	}
}
