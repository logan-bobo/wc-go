package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func countLines(bytes []byte) int {
	var lineCount int

	for _, byte := range bytes {
		// The decimal representation of a UTF-8 newline is 10.
		if byte == 10 {
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
	lines := strings.SplitSeq(fileContent, "\n")

	for line := range lines {
		arrLine := strings.SplitSeq(line, "")
		for range arrLine {
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

	words := strings.SplitSeq(text, " ")
	for word := range words {
		tabInWord := strings.Contains(word, "	")
		if tabInWord {
			splitWords := strings.SplitSeq(word, "	")

			for splitWord := range splitWords {
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
	text := strings.SplitSeq(fileContent, "\n")

	for line := range text {
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
	var fileBytes []byte
	var fileName string

	bytesFlag := flag.Bool("c", false, "Count the bytes in a file/stdin")
	linesFlag := flag.Bool("l", false, "Count the number of lines in a file/stdin")
	wordsFlag := flag.Bool("w", false, "Count the number of words in a file/stdin")
	charsFlag := flag.Bool("m", false, "Count the number of characters in a file/stdin")
	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		fileName = args[0]
		fileBytes = fileToBytes(fileName)

	} else {
		stdin, err := io.ReadAll(os.Stdin)

		if err != nil {
			panic(err)
		}

		fileBytes = stdin
	}

	if *bytesFlag {
		fileByteTotal := countBytes(fileBytes)
		fmt.Println(fileByteTotal, fileName)
	}

	if *linesFlag {
		fileLines := countLines(fileBytes)
		fmt.Println(fileLines, fileName)
	}

	if *wordsFlag {
		fileWords := countWords(fileBytes)
		fmt.Println(fileWords, fileName)
	}

	if *charsFlag {
		fileChars := countChars(fileBytes)
		fmt.Println(fileChars, fileName)
	}

	if !*bytesFlag && !*linesFlag && !*wordsFlag && !*charsFlag {
		fileByteCount := countBytes(fileBytes)
		fileLines := countLines(fileBytes)
		fileWords := countWords(fileBytes)
		fmt.Println(fileLines, fileWords, fileByteCount, fileName)
	}
}
