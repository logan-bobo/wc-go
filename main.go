package main

import (
	"flag"
	"fmt"

	"github.com/logan-bobo/wc-go/cmd"
)

func main() {
	var fileName string
	bytesFlag := flag.String("c", "", "Count the bytes in a file")
	linesFlag := flag.String("l", "", "Count the number of lines in a file")
	wordsFlag := flag.String("w", "", "Count the number of words in a file")
	charFlag := flag.String("m", "", "Count the number of characters in a file")
	flag.Parse()

	if len(*bytesFlag) > 0 {
		fileName = *bytesFlag
		fileBytes := cmd.CountBytes(fileName)
		fmt.Println(" ", fileBytes, fileName)
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
		fmt.Println(fileName)
		fileBytes := cmd.CountBytes(fileName)
		fileLines := cmd.CountLines(fileName)
		fileWords := cmd.CountWords(fileName)
		fmt.Println(" ", fileBytes, fileLines, fileWords, fileName)
	}
}
