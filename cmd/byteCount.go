package cmd

import (
	"os"
)

func CountBytes(fileName string) int64 {
	fileData, err := os.Stat(fileName)

	if err != nil {
		panic(err)
	}

	fileSize := fileData.Size()

	return fileSize
}
