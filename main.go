package main

import (
	"flag"
	"fmt"

	"github.com/logan-bobo/wc-go/cmd"
)

func main() {
	var fileName string
	bytesFlag := flag.String("c", "", "Count the bytes from a file or stdin")
	linesFlag := flag.String("l", "", "Count the number of lines in a file or stdin")
	wordsFlag := flag.String("w", "", "Count the number of words in a file or stdin")
	charFlag := flag.String("m", "", "Count the number of characters in a file or stdin")
	flag.Parse()

	if len(*bytesFlag) > 0 {
		fileName = *bytesFlag
		fileBytes := cmd.FileToBytes(fileName)
		fileByteTotal := cmd.CountBytes(fileBytes)
		fmt.Println(fileByteTotal, fileName)

	} else if len(*linesFlag) > 0 {
		fileName = *linesFlag
		fileBytes := cmd.FileToBytes(fileName)
		fileLines := cmd.CountLines(fileBytes)
		fmt.Println(fileLines, fileName)

	} else if len(*wordsFlag) > 0 {
		fileName = *wordsFlag
		fileBytes := cmd.FileToBytes(fileName)
		fileWords := cmd.CountWords(fileBytes)
		fmt.Println(fileWords, fileName)

	} else if len(*charFlag) > 0 {
		fileName = *charFlag
		fileBytes := cmd.FileToBytes(fileName)
		fileChars := cmd.CountChars(fileBytes)
		fmt.Println(fileChars, fileName)

	} else {
		fileName := flag.Arg(0)
		fileBytes := cmd.FileToBytes(fileName)
		fileByteCount := cmd.CountBytes(fileBytes)
		fileLines := cmd.CountLines(fileBytes)
		fileWords := cmd.CountWords(fileBytes)
		fmt.Println(fileLines, fileWords, fileByteCount, fileName)
	}
}
