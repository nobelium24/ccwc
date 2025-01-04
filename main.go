package main

import (
	countbyte "ccwc/countByte"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Print("Incomplete command line args")
	}
	filePath := os.Args[1]
	operation := os.Args[2]
	if operation == "-c" {
		bytes, err := countbyte.CountByte(filePath)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
		}
		fmt.Print(bytes, filePath)
	}
}
