package c3

import (
	"log"
	"os"
)

func ReName() {
	sourcePath := "c3/test.txt"
	newPath := "c3/text.txt"
	err := os.Rename(sourcePath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
