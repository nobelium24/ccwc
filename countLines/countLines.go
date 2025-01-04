package countlines

import (
	"bufio"
	"fmt"
	"os"
)

func CountLines(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()
	numLines := 0
	input := bufio.NewScanner(file)
	for input.Scan() {
		numLines++
	}
	if err := input.Err(); err != nil {
		return 0, fmt.Errorf("Error reading file: %v", err)
	}
	return numLines, nil
}
