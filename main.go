package main

import (
	charactercount "ccwc/characterCount"
	countbyte "ccwc/countByte"
	countlines "ccwc/countLines"
	countwords "ccwc/countWords"
	"fmt"
	"io"
	"os"
)

func main() {
	var input io.Reader
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file_path> <operation>")
		return
	}

	if len(os.Args) == 2 {
		input = os.Stdin
	} else {
		filePath := os.Args[1]
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
		}
		defer file.Close()
		input = file
	}
	operation := os.Args[len(os.Args)-1]
	if operation == "-c" {
		bytes, err := countbyte.CountByte(input)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
		}
		fmt.Print(bytes)
	} else if operation == "-l" {
		lines, err := countlines.CountLines(input)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
		}
		fmt.Print(lines)
	} else if operation == "-w" {
		words, err := countwords.CountWord(input)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
		}
		fmt.Print(words)
	} else if operation == "-m" {
		characters, err := charactercount.CharacterCount(input)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
		}
		fmt.Print(characters)
	} else if operation == "" {
		bytes, err := countbyte.CountByte(input)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
		}
		lines, err := countlines.CountLines(input)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
		}
		words, err := countwords.CountWord(input)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
		}
		fmt.Print(bytes, lines, words)
	}
}
