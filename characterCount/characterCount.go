package charactercount

import (
	"bufio"
	"fmt"
	"io"
	// "os"
)

func CharacterCount(filePath io.Reader) (int, error) {
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	return 0, fmt.Errorf("Error accessing file: %v", err)
	// }
	// defer file.Close()

	input := bufio.NewReader(filePath)
	characterCount := 0
	byteSlice := make([]byte, 4096)
	for {
		n, err := input.Read(byteSlice)
		if err != nil {
			return 0, fmt.Errorf("Error reading characters: %v", err)
		}
		if n == 0 {
			break
		}
		characterCount += n
	}
	return characterCount, nil
}
