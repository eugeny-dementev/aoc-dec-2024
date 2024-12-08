package main

import (
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var example string = `@~do()mul(683,461) >,~select()what()};<mul(848,589)don't()2397dk?sd.,.*mul(848,589)`

func readDay3Input() []byte {
	content, err := os.ReadFile("day3.txt")
	assert.NoError(err, "should read day3.txt file without issues")

	return content
}

func day3ExtractMul() {
	content := readFileInput("day3.txt")

	matchSearcher := regexp.MustCompile(`mul\([[:digit:]]+,[[:digit:]]+\)`)
	matches := matchSearcher.FindAll(content, -1)

	numbersSearcher := regexp.MustCompile(`[[:digit:]]+`)

	var result int64 = 0

	for _, match := range matches {
		numbers := numbersSearcher.FindAll(match, -1)

		result += day3CalculateMul(numbers)
	}

	fmt.Println("Sum of muls is:", result)
}

func day3ExtractMulWithCondition() {
	content := readFileInput("day3.txt")

	matchSearcher := regexp.MustCompile(`mul\([[:digit:]]+,[[:digit:]]+\)`)
	matches := matchSearcher.FindAll(content, -1)
	parts := matchSearcher.Split(string(content), -1)

	operationActive := true

	numbersSearcher := regexp.MustCompile(`[[:digit:]]+`)

	var result int64 = 0

	for i, match := range matches {
		if strings.Contains(parts[i], "don't()") {
			operationActive = false
		} else if strings.Contains(parts[i], "do()") {
			operationActive = true
		}

		if operationActive {
			numbers := numbersSearcher.FindAll(match, -1)
			result += day3CalculateMul(numbers)
		}
	}

	fmt.Println("Sum of muls is:", result)
}

func day3CalculateMul(numbers [][]byte) int64 {
	leftInt, err := strconv.ParseInt(string(numbers[0]), 0, 0)
	assert.NoError(err, "should ParseInt without any issues", "numbers[0]", numbers[0])
	rightInt, err := strconv.ParseInt(string(numbers[1]), 0, 0)
	assert.NoError(err, "should ParseInt without any issues", "numbers[1]", numbers[1])

	return leftInt * rightInt
}
