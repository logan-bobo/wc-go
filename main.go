package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func stripTabAndSpaceFromLine(line string) []string {
	var flatWords []string

	words := strings.Split(line, " ")
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

func countBytes(fileName string) int64 {
	fileData, err := os.Stat(fileName)
	check(err)
	fileSize := fileData.Size()

	return fileSize
}

func countLines(fileName string) int {
	var lineCount int

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}

func countWords(fileName string) int {
	var wordCount int
	var lineCount int

	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line != " " {

			words := stripTabAndSpaceFromLine(line)
			if words[0] != "" {
				wordCount += len(words)
				lineCount++
			}
		}
	}

	return wordCount
}

func main() {
	var fileName string

	// Build out flags
	bytesFlag := flag.String("c", "", "Count the bytes in a file")
	linesFlag := flag.String("l", "", "Count the number of lines in a file")
	wordsFlag := flag.String("w", "", "Count the number of words in a file")

	flag.Parse()

	// Return bytes in a file
	if len(*bytesFlag) > 0 {
		fileName = *bytesFlag
		fileBytes := countBytes(fileName)
		fmt.Println(" ", fileBytes, fileName)
	}

	// Return lines in a file
	if len(*linesFlag) > 0 {
		fileName = *linesFlag
		fileLines := countLines(fileName)
		fmt.Println(" ", fileLines, fileName)
	}

	// Return the words in a file
	if len(*wordsFlag) > 0 {
		fileName = *wordsFlag
		fileWords := countWords(fileName)
		fmt.Println(" ", fileWords, fileName)
	}
}
