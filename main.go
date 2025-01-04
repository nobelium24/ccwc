package main

import (
	countbyte "ccwc/countByte"
	countlines "ccwc/countLines"
	countwords "ccwc/countWords"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Print("Incomplete command line args \n")
	}
	filePath := os.Args[1]
	operation := os.Args[2]
	if operation == "-c" {
		bytes, err := countbyte.CountByte(filePath)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
		}
		fmt.Print(bytes, filePath)
	} else if operation == "-l" {
		lines, err := countlines.CountLines(filePath)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
		}
		fmt.Print(lines, filePath)
	} else if operation == "-w" {
		words, err := countwords.CountWords(filePath)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
		}
		fmt.Print(words, filePath)
	}
}
