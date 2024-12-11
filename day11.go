package main

import (
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
	"strconv"
	"strings"
)

var day11example = `125 17`

func splitRock(rocks []string, i int) []string {
	rock := rocks[i]
	length := len(rock)

	assert.Assert(length%2 == 0, "rock digits length should be even", "rock", rock, "length", length)

	middle := length / 2
	left, right := rock[:middle], rock[middle:]

	right = strings.TrimLeft(right, "0")

  if len(right) == 0 {
    right = "0"
  }

	return append(rocks[:i], append([]string{left, right}, rocks[i+1:]...)...)
}

func day11Rocks() {
	line := readFileLines("day11.txt")[0]
	// line := day11example

	fmt.Println("line", line)

	rocks := strings.Split(line, " ")

	fmt.Println("Parsed rocks:", rocks)

  for blink := range 25 {
		for i := 0; i < len(rocks); i++ {
			if rocks[i] == "0" {
				rocks[i] = "1"
			} else if len(rocks[i])%2 == 0 {
				rocks = splitRock(rocks, i)
        i++
			} else {
				num, err := strconv.ParseInt(rocks[i], 0, 0)
				assert.NoError(err, "should parse rock number with no err", "rocks[i]", rocks[i])

				num *= 2024
				rocks[i] = strconv.FormatInt(num, 10)
			}
		}
    fmt.Println("Blink number", blink)
	}

  fmt.Println("Stones after 25 blinks", len(rocks))
}
