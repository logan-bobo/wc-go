package main

import (
	"os"
	"reflect"
	"testing"
)

const testFile string = "test.txt"

func TestFileToBytes(t *testing.T) {
	fileData := FileToBytes(testFile)
	fileBytes, err := os.ReadFile(testFile)

	if err != nil {
		t.Error("Can not find test data")
	}

	equalBytes := reflect.DeepEqual(fileData, fileBytes)

	if equalBytes != true {
		t.Errorf("FileToBytes(%s), FAILED. Expected %s got %s \n", testFile, fileBytes, fileData)
	} else {
		t.Logf("FileToBytes(%s), PASSED.\n", testFile)
	}
}

func TestByteCount(t *testing.T) {
	fileBytes := FileToBytes(testFile)
	byteCount := CountBytes(fileBytes)
	byteCountExpected := 342190

	if byteCount != byteCountExpected {
		t.Errorf("FileToBytes(%s), FAILED. Expected %s got %d \n", testFile, fileBytes, len(fileBytes))
	} else {
		t.Logf("FileToBytes(%s), PASSED.\n", testFile)
	}
}

func TestLineCount(t *testing.T) {
	fileBytes := FileToBytes(testFile)
	lineCount := CountLines(fileBytes)
	lineCountExpected := 7145

	if lineCount != lineCountExpected {
		t.Errorf("FileToBytes(%s), FAILED. Expected %s got %d \n", testFile, fileBytes, len(fileBytes))
	} else {
		t.Logf("FileToBytes(%s), PASSED.\n", testFile)
	}
}
