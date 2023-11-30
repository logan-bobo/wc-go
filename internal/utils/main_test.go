package utils

import (
	"os"
	"reflect"
	"testing"
)

func TestFileToBytes(t *testing.T) {
	fileName := "../../test.txt"
	fileData := FileToBytes(fileName)
	fileBytes, err := os.ReadFile(fileName)

	if err != nil {
		t.Error("Can not find test data")
	}

	equalBytes := reflect.DeepEqual(fileData, fileBytes)

	if equalBytes != true {
		t.Errorf("FileToBytes(%s), FAILED. Expected %s got %s \n", fileName, fileBytes, fileData)
	} else {
		t.Logf("FileToBytes(%s), PASSED.\n", fileName)
	}
}