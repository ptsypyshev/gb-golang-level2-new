package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

const readOnlyFilePath = "ro.txt"

func main() {
	f, err := os.Create(readOnlyFilePath)
	if err != nil {
		log.Fatalf("cannot create file %s: %s\n", readOnlyFilePath, err)
	}
	defer f.Close()
	_, err = f.WriteString("first line")
	if err != nil {
		log.Fatalf("cannot write to file %s: %s\n", readOnlyFilePath, err)
	}

	err = f.Chmod(fs.FileMode(os.O_RDONLY))
	if err != nil {
		log.Fatalf("cannot chmod file %s: %s\n", readOnlyFilePath, err)
	}
	f.Close()

	f, err = os.Open(readOnlyFilePath)
	if err != nil {
		log.Fatalf("cannot open file %s: %s\n", readOnlyFilePath, err)
	}

	_, err = f.WriteString("test string")
	if err != nil {
		log.Printf("cannot update file %s: %s\n", readOnlyFilePath, err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		log.Fatalf("cannot get filestat %s: %s\n", readOnlyFilePath, err)
	}

	buf := make([]byte, fileInfo.Size())
	_, err = f.Read(buf)
	if err != nil {
		log.Fatalf("cannot read file %s: %s\n", readOnlyFilePath, err)
	}

	fmt.Println("file content:")
	fmt.Printf("%s", buf)
}
