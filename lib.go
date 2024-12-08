package main

import (
	"bufio"
	"bytes"
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
	"os"
	"strings"
)

func readInput(fileName string) []byte {
	content, err := os.ReadFile(fileName)
	assert.NoError(err, fmt.Sprintf("should read %s file without issues", fileName))

	return bytes.TrimSpace(content)
}

func readLines(fileName string) []string {
	input := readInput(fileName)

	var lines []string

	sc := bufio.NewScanner(strings.NewReader(string(input)))

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

func readMap(fileName string) [][]string {
	lines := readLines(fileName)

	var board [][]string

	for _, line := range lines {
		points := strings.Split(line, "")
		board = append(board, points)
	}

	return board
}
