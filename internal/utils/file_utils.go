package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ReadFileContent(relativePath string) (string, error) {

	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %w", err)
	}

	path := filepath.Join(wd, relativePath)

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
