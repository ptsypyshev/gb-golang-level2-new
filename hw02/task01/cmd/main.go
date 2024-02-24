package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	defaultFileName = "notes.txt"
	exitCmd         = "exit"
)

var ErrFileWrite = errors.New("cannot create/write file")

func main() {
	var lineCount int

	f, err := os.Create(defaultFileName)
	if err != nil {
		log.Fatalf("cannod read from stdin: %s\n", err)
	}
	defer f.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your note or 'exit' to close app: ")
		lineValue, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("cannod read from stdin: %s\n", err)
			return
		}
		lineValue = strings.TrimSpace(lineValue)

		if lineValue == exitCmd {
			fmt.Println("Bye")
			return
		}
		lineCount++

		f.WriteString(fmt.Sprintf("%d %s %s\n", lineCount, time.Now().Format(time.DateTime), lineValue))
	}
}
