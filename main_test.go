package main

import (
	"os"
	"reflect"
	"testing"
)

const testFile string = "test.txt"

func TestFileToBytes(t *testing.T) {
	fileData := fileToBytes(testFile)
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
	fileBytes := fileToBytes(testFile)
	byteCount := countBytes(fileBytes)
	byteCountExpected := 342190

	if byteCount != byteCountExpected {
		t.Errorf("FileToBytes(%s), FAILED. Expected %s got %d \n", testFile, fileBytes, len(fileBytes))
	} else {
		t.Logf("FileToBytes(%s), PASSED.\n", testFile)
	}
}

func TestLineCount(t *testing.T) {
	fileBytes := fileToBytes(testFile)
	lineCount := countLines(fileBytes)
	lineCountExpected := 7145

	if lineCount != lineCountExpected {
		t.Errorf("CountLines(%s), FAILED. Expected %s got %d \n", testFile, fileBytes, len(fileBytes))
	} else {
		t.Logf("CountLines(%s), PASSED.\n", testFile)
	}
}

func TestWordCount(t *testing.T) {
	fileBytes := fileToBytes(testFile)
	wordCount := countWords(fileBytes)
	wordCountExpected := 58164

	if wordCount != wordCountExpected {
		t.Errorf("CountWords(%s), FAILED. Expected %s got %d \n", testFile, fileBytes, len(fileBytes))
	} else {
		t.Logf("CountWords(%s), PASSED.\n", testFile)
	}
}

func TestCharCount(t *testing.T) {
	fileBytes := fileToBytes(testFile)
	charCount := countChars(fileBytes)
	charCountExpected := 339292

	if charCount != charCountExpected {
		t.Errorf("CountChars(%s), FAILED. Expected %s got %d \n", testFile, fileBytes, charCountExpected)
	} else {
		t.Logf("CountChars(%s), PASSED.\n", testFile)
	}
}
