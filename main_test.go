package main

import (
	"os"
	"reflect"
	"testing"
)

const testFile string = "test.txt"

func TestByteCount(t *testing.T) {
	fileBytes, _ := os.ReadFile(testFile)
	byteCount := countBytes(fileBytes)
	byteCountExpected := 342190

	if byteCount != byteCountExpected {
		t.Errorf("countBytes(%s), FAILED. Expected %d got %d \n", testFile, byteCountExpected, byteCount)
	} else {
		t.Logf("countBytes(%s), PASSED.\n", testFile)
	}
}

func TestLineCount(t *testing.T) {
	fileBytes, _ := os.ReadFile(testFile)
	lineCount := countLines(fileBytes)
	lineCountExpected := 7145

	if lineCount != lineCountExpected {
		t.Errorf("countLines(%s), FAILED. Expected %d got %d \n", testFile, lineCountExpected, lineCount)
	} else {
		t.Logf("countLines(%s), PASSED.\n", testFile)
	}
}

func TestWordCount(t *testing.T) {
	fileBytes, _ := os.ReadFile(testFile)
	wordCount := countWords(fileBytes)
	wordCountExpected := 58164

	if wordCount != wordCountExpected {
		t.Errorf("countWords(%s), FAILED. Expected %d got %d \n", testFile, wordCountExpected, wordCount)
	} else {
		t.Logf("countWords(%s), PASSED.\n", testFile)
	}
}

func TestCharCount(t *testing.T) {
	fileBytes, _ := os.ReadFile(testFile)
	charCount := countChars(fileBytes)
	charCountExpected := 339292

	if charCount != charCountExpected {
		t.Errorf("countChars(%s), FAILED. Expected %d got %d \n", testFile, charCountExpected, charCount)
	} else {
		t.Logf("countChars(%s), PASSED.\n", testFile)
	}
}

func TestStripTabandSpaceFromLine(t *testing.T) {
	line := "hello	world.	today."
	outputSlice := []string{"hello", "world.", "today."}

	stripLine := stripTabAndSpaceFromLine(line)

	isEqual := reflect.DeepEqual(outputSlice, stripLine)

	if isEqual != true {
		t.Errorf("stripTabAndSpaceFromLine(%s), FAILED. Expected %v got %v \n", line, outputSlice, stripLine)
	} else {
		t.Logf("stripTabAndSpaceFromLine(%s), PASSED.\n", line)
	}
}
