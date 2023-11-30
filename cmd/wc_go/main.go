package main

import (
	"fmt"
	"io"
	"os"

	"github.com/logan-bobo/wc-go/internal/utils"
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
			fileBytes = utils.FileToBytes(fileName)
		} else {
			for _, legalFlag := range legalFlags {
				if os.Args[1] == legalFlag {
					flag = os.Args[1]
					fileName = os.Args[2]
					fileBytes = utils.FileToBytes(fileName)
				}
			}
		}
	}

	if flag == "-c" {
		fileByteTotal := utils.CountBytes(fileBytes)
		fmt.Println(fileByteTotal, fileName)

	} else if flag == "-l" {
		fileLines := utils.CountLines(fileBytes)
		fmt.Println(fileLines, fileName)

	} else if flag == "-w" {
		fileWords := utils.CountWords(fileBytes)
		fmt.Println(fileWords, fileName)

	} else if flag == "-m" {
		fileChars := utils.CountChars(fileBytes)
		fmt.Println(fileChars, fileName)

	} else {
		fileByteCount := utils.CountBytes(fileBytes)
		fileLines := utils.CountLines(fileBytes)
		fileWords := utils.CountWords(fileBytes)
		fmt.Println(fileLines, fileWords, fileByteCount, fileName)
	}
}
