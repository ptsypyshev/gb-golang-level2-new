package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"
)

const (
	defaultFileName = "notes.txt"
	exitCmd         = "exit"
)

func main() {
	var (
		lineCount int
		buf       bytes.Buffer
	)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your note or 'exit' to close app: ")
		lineValue, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("cannod read from stdin: %s\n", err)
			break
		}
		lineValue = strings.TrimSpace(lineValue)

		if lineValue == exitCmd {
			fmt.Println("Bye")
			break
		}
		lineCount++

		buf.WriteString(fmt.Sprintf("%d %s %s\n", lineCount, time.Now().Format(time.DateTime), lineValue))
	}

	// ioutil.WriteFile() is deprecated: As of Go 1.16, this function simply calls [os.WriteFile]
	err := os.WriteFile(defaultFileName, buf.Bytes(), fs.ModePerm)
	if err != nil {
		log.Fatalf("cannot write to file %s: %s\n", defaultFileName, err)
	}

	// ioutil.ReadFile()  is deprecated: As of Go 1.16, this function simply calls os.ReadFile.
	data, err := os.ReadFile(defaultFileName)
	if err != nil {
		log.Fatalf("cannot read file %s: %s\n", defaultFileName, err)
	}
	fmt.Println(string(data))
}
