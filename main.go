package main

import (
	"flag"
	"fmt"
	"os"
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

func main () {
	// Build out flags
	bytesFlag := flag.String("c", "", "Count the bytes in a file")
	flag.Parse()
	fileName := *bytesFlag

	fileBytes := countBytes(fileName)
	fmt.Println(" ", fileBytes, fileName)
}