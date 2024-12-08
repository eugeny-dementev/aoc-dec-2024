package main

import (
	"bufio"
	"bytes"
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
	"os"
	"strings"
)

func readFileInput(fileName string) []byte {
	content, err := os.ReadFile(fileName)
	assert.NoError(err, fmt.Sprintf("should read %s file without issues", fileName))

	return bytes.TrimSpace(content)
}

func readFileLines(fileName string) []string {
	input := readFileInput(fileName)

	var lines []string

	sc := bufio.NewScanner(strings.NewReader(string(input)))

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

func readFileMap(fileName string) [][]string {
	lines := readFileLines(fileName)

	var board [][]string

	for _, line := range lines {
		points := strings.Split(line, "")
		board = append(board, points)
	}

	return board
}
