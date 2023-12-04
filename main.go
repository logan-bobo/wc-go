package main

import (
	"fmt"
	"io"
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
				if os.Args[1] == legalFlag {
					panic("Flag set with no filename")
				}
			}
			fileName = os.Args[1]
			fileBytes = FileToBytes(fileName)
		} else {
			for _, legalFlag := range legalFlags {
				if os.Args[1] == legalFlag {
					flag = os.Args[1]
					fileName = os.Args[2]
					fileBytes = FileToBytes(fileName)
				}
			}
		}
	}

	if flag == "-c" {
		fileByteTotal := CountBytes(fileBytes)
		fmt.Println(fileByteTotal, fileName)

	} else if flag == "-l" {
		fileLines := CountLines(fileBytes)
		fmt.Println(fileLines, fileName)

	} else if flag == "-w" {
		fileWords := CountWords(fileBytes)
		fmt.Println(fileWords, fileName)

	} else if flag == "-m" {
		fileChars := CountChars(fileBytes)
		fmt.Println(fileChars, fileName)

	} else {
		fileByteCount := CountBytes(fileBytes)
		fileLines := CountLines(fileBytes)
		fileWords := CountWords(fileBytes)
		fmt.Println(fileLines, fileWords, fileByteCount, fileName)
	}
}
