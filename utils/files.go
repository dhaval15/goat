package utils

import (
	// "bufio"
	"os"
	// "path/filepath"
)

func SaveFileToDirectory(fileBytes []byte, directory, filename string) error {
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return err
	}

	// filePath := filepath.Join(directory, filename)
	//file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	//bufio.NewWriter(file)
	return nil
	//return writer.WriteByte(filePath, fileBytes, 0644)
}
