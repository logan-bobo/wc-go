package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
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
		lineCount ++
	}
	return lineCount
}

func main () {
	var fileName string
	
	// Build out flags
	bytesFlag := flag.String("c", "", "Count the bytes in a file")
	linesFlag := flag.String("l", "", "Count the number of lines in a file")
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
}