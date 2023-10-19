package main

import (
	"flag"
	"fmt"

	"github.com/logan-bobo/wc-go/cmd"
)

func main() {
	var fileName string

	// Build out flags
	bytesFlag := flag.String("c", "", "Count the bytes in a file")
	linesFlag := flag.String("l", "", "Count the number of lines in a file")
	wordsFlag := flag.String("w", "", "Count the number of words in a file")
	charFlag := flag.String("m", "", "Count the number of characters in a file")
	flag.Parse()

	// Return bytes in a file
	if len(*bytesFlag) > 0 {
		fileName = *bytesFlag
		fileBytes := cmd.CountBytes(fileName)
		fmt.Println(" ", fileBytes, fileName)
	}

	// Return lines in a file
	if len(*linesFlag) > 0 {
		fileName = *linesFlag
		fileLines := cmd.CountLines(fileName)
		fmt.Println(" ", fileLines, fileName)
	}

	// Return the words in a file
	if len(*wordsFlag) > 0 {
		fileName = *wordsFlag
		fileWords := cmd.CountWords(fileName)
		fmt.Println(" ", fileWords, fileName)
	}

	// Return the words in a file
	if len(*charFlag) > 0 {
		fileName = *charFlag
		fileWords := cmd.CountChars(fileName)
		fmt.Println(" ", fileWords, fileName)
	}
}
