package countwords

import (
	"bufio"
	"fmt"
	"os"
)

func CountWords(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("Error opening file:%v", err)
	}
	defer file.Close()

	var wordCount int
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordCount++
	}
	if err := input.Err(); err != nil {
		return 0, fmt.Errorf("Error reading file: %v", err)
	}
	return wordCount, nil
}
