package main

import (
	"fmt"
	"io"
	"os"

	"github.com/logan-bobo/wc-go/internal/wc_utils"
)

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
			fileBytes = wcutils.FileToBytes(fileName)
		} else {
			for _, legalFlag := range legalFlags {
				if os.Args[1] == legalFlag {
					flag = os.Args[1]
					fileName = os.Args[2]
					fileBytes = wcutils.FileToBytes(fileName)
				}
			}
		}
	}

	if flag == "-c" {
		fileByteTotal := wcutils.CountBytes(fileBytes)
		fmt.Println(fileByteTotal, fileName)

	} else if flag == "-l" {
		fileLines := wcutils.CountLines(fileBytes)
		fmt.Println(fileLines, fileName)

	} else if flag == "-w" {
		fileWords := wcutils.CountWords(fileBytes)
		fmt.Println(fileWords, fileName)

	} else if flag == "-m" {
		fileChars := wcutils.CountChars(fileBytes)
		fmt.Println(fileChars, fileName)

	} else {
		fileByteCount := wcutils.CountBytes(fileBytes)
		fileLines := wcutils.CountLines(fileBytes)
		fileWords := wcutils.CountWords(fileBytes)
		fmt.Println(fileLines, fileWords, fileByteCount, fileName)
	}
}
