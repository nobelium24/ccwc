package countlines

import (
	"bufio"
	"fmt"
	"io"
	// "os"
)

func CountLines(filePath io.Reader) (int, error) {
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	return 0, fmt.Errorf("Error opening file: %v", err)
	// }
	// defer file.Close()
	numLines := 0
	input := bufio.NewScanner(filePath)
	for input.Scan() {
		numLines++
	}
	if err := input.Err(); err != nil {
		return 0, fmt.Errorf("Error reading file: %v", err)
	}
	return numLines, nil
}
