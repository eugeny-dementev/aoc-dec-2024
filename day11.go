package main

import (
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
	"strconv"
	"strings"
)

var day11example = `125 17`

func splitRock(rock string) (left, right string) {
	length := len(rock)

	assert.Assert(length%2 == 0, "rock digits length should be even", "rock", rock, "length", length)

	middle := length / 2
	left, right = rock[:middle], rock[middle:]

	right = strings.TrimLeft(right, "0")

	if len(right) == 0 {
		right = "0"
	}

	return left, right
}

type RocksParams struct {
	rock  string
	depth int
}

var rocksCache map[RocksParams]int64 = map[RocksParams]int64{}

func countProgressionSum(elem string, depth, targetDepth int) int64 {
	if depth == targetDepth {
		return 1
	}

	key := RocksParams{elem, depth}

	result, ok := rocksCache[key]
	if ok {
		return result
	}

	if elem == "0" {
		result = countProgressionSum("1", depth+1, targetDepth)
	} else if len(elem)%2 == 0 {
		left, right := splitRock(elem)
		leftRes, rightRes := countProgressionSum(left, depth+1, targetDepth), countProgressionSum(right, depth+1, targetDepth)
		// fmt.Println("Split", elem, "into", left, right)
		result = leftRes + rightRes
	} else {

		num, err := strconv.ParseInt(elem, 0, 0)
		assert.NoError(err, "should parse rock number with no err", "elem", elem)

		num *= 2024
		rock := strconv.FormatInt(num, 10)

		result = countProgressionSum(rock, depth+1, targetDepth)

	}

	rocksCache[key] = result

	return result
}

func day11Recursion() {
	line := readFileLines("day11.txt")[0]
	// line := day11example
	// line := "0"

	rocks := strings.Split(line, " ")

	fmt.Println("Parsed rocks:", rocks)

	var totalRocks int64 = 0

	for _, rock := range rocks {
		totalRocks += countProgressionSum(rock, 0, 75)
	}

	// fmt.Println("Rocks cache", rocksCache)

	fmt.Println("Total rocks:", totalRocks)
}
