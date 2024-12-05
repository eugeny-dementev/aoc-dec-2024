package main

import (
	"bytes"
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
	"os"
)

func readInput(fileName string) []byte {
	content, err := os.ReadFile(fileName)
	assert.NoError(err, fmt.Sprintf("should read %s file without issues", fileName))

	return bytes.TrimSpace(content)
}
