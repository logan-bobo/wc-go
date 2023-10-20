package main

import (
	"flag"
	"fmt"

	"github.com/logan-bobo/wc-go/cmd"
)

func main() {
	var fileName string
	bytesFlag := flag.String("c", "", "Count the bytes from a file or stdin")
	linesFlag := flag.String("l", "", "Count the number of lines in a file")
	wordsFlag := flag.String("w", "", "Count the number of words in a file")
	charFlag := flag.String("m", "", "Count the number of characters in a file")
	flag.Parse()

	if len(*bytesFlag) > 0 {
		fileName = *bytesFlag
		if fileName == "stdin" {
			fmt.Println("stdin")
		}

		fileBytes := cmd.FileToBytes(fileName)
		fileByteTotal := cmd.CountBytes(fileBytes)
		fmt.Println(" ", fileByteTotal, fileName)

	} else if len(*linesFlag) > 0 {
		fileName = *linesFlag
		fileLines := cmd.CountLines(fileName)
		fmt.Println(" ", fileLines, fileName)

	} else if len(*wordsFlag) > 0 {
		fileName = *wordsFlag
		fileWords := cmd.CountWords(fileName)
		fmt.Println(" ", fileWords, fileName)

	} else if len(*charFlag) > 0 {
		fileName = *charFlag
		fileChars := cmd.CountChars(fileName)
		fmt.Println(" ", fileChars, fileName)

	} else {
		fileName := flag.Arg(0)
		fileBytes := cmd.FileToBytes(fileName)
		fileByteCount := cmd.CountBytes(fileBytes)
		fileLines := cmd.CountLines(fileName)
		fileWords := cmd.CountWords(fileName)
		fmt.Println(" ", fileByteCount, fileLines, fileWords, fileName)
	}
}
