package utils

import (
	"log"
	"os"
)

func GetFileReader(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		errMessage := err.Error()
		log.Fatalf("Unable to open the file %s. Error Message: %s\n", path, errMessage)
	}

	return file
}
