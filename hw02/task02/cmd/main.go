package main

import (
	"fmt"
	"log"
	"os"
)

const (
	defaultFileName = "notes.txt"
)

func main() {
	fileInfo, err := os.Stat(defaultFileName)
	if err != nil || fileInfo.Size() == 0 {
		log.Fatalf("file %s is empty or absent", defaultFileName)
	}

	f, err := os.OpenFile(defaultFileName, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("cannod read from file: %s\n", err)
	}
	defer f.Close()

	buf := make([]byte, fileInfo.Size())
	f.Read(buf)
	fmt.Printf("%s", buf)
}
