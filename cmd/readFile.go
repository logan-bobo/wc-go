package cmd

import (
	"os"
)

func FileToBytes(fileName string) []byte {
	bytes, err := os.ReadFile(fileName)
	
	if err != nil {
		panic(err)
	}

	return bytes
}
