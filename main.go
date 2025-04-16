package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type metadata struct {
	bytes int
	lines int
	words int
	chars int
}

func newMetadata() *metadata {
	return &metadata{}
}

func (m *metadata) populate(bytes []byte) {
	byteCount := countBytes(bytes)
	lineCount := countLines(bytes)
	wordCount := countWords(bytes)
	charCount := countWords(bytes)

	m.bytes += byteCount
	m.lines += lineCount
	m.words += wordCount
	m.chars += charCount
}

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

func populateMetadataFromReader(reader bufio.Reader, fileMetadata *metadata) {
	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fileMetadata.populate(line)
	}
}

func main() {
	var fileName string

	fileMetadata := newMetadata()

	bytesFlag := flag.Bool("c", false, "Count the bytes in a file/stdin")
	linesFlag := flag.Bool("l", false, "Count the number of lines in a file/stdin")
	wordsFlag := flag.Bool("w", false, "Count the number of words in a file/stdin")
	charsFlag := flag.Bool("m", false, "Count the number of characters in a file/stdin")
	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		fileName = args[0]
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		populateMetadataFromReader(*reader, fileMetadata)

	} else {
		reader := bufio.NewReader(os.Stdin)
		populateMetadataFromReader(*reader, fileMetadata)
	}

	if *bytesFlag {
		fmt.Println(fileMetadata.bytes, fileName)
	}

	if *linesFlag {
		fmt.Println(fileMetadata.lines, fileName)
	}

	if *wordsFlag {
		fmt.Println(fileMetadata.words, fileName)
	}

	if *charsFlag {
		fmt.Println(fileMetadata.chars, fileName)
	}

	if !*bytesFlag && !*linesFlag && !*wordsFlag && !*charsFlag {
		fmt.Println(fileMetadata.lines, fileMetadata.words, fileMetadata.bytes, fileName)
	}
}
