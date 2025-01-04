package countbyte

import (
	"fmt"
	"io"
	"os"
)

func CountByte(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("Error opening file:%v", err)
	}
	defer file.Close()
	buffer := make([]byte, 4096)
	// var fileData []byte
	var totalBytes int
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return 0, fmt.Errorf("error reading file: %w", err)
		}
		if n == 0 {
			break
		}
		totalBytes += n
		// fileData = append(fileData, buffer[:n]...)
	}
	return totalBytes, nil
}
