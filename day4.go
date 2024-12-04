package main

import (
	"bytes"
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
)

var target = "XMAS"

var (
	X byte = []byte("X")[0]
	M byte = []byte("M")[0]
	A byte = []byte("A")[0]
	S byte = []byte("S")[0]
)

var day4example1 []byte = []byte(`..X...
.SAMX.
.A..A.
XMAS.S
.X....`)

var day4example2 []byte = []byte(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

type Matrix [][]byte

func readMatrix(content []byte) Matrix {
	lines := bytes.Split(bytes.TrimSpace(content), []byte("\n"))

	matrix := make(Matrix, len(lines))

	for l, line := range lines {
		letters := bytes.Split(bytes.TrimSpace(line), []byte(""))

		chars := make([]byte, len(letters))

		for i, char := range letters {
			assert.Assert(len(char) == 1, "len of each char in split line should be 1", "char", char)

			chars[i] = char[0]
		}

		matrix[l] = chars
	}

	return matrix
}

func (m Matrix) Print() {
	for _, line := range m {
		for _, char := range line {
			fmt.Printf("%s", string(char))
		}
		fmt.Printf("\n")
	}
}

func day4IsTargetWork(word []byte) int {
	str := string(word)
  isTarget := str == target

	fmt.Println("WORD:", str, target, isTarget)

  if isTarget {
    return 1
  }

	return 0
}

func day4CheckCoordinate(x, y int, matrix Matrix, direction int) int {
	assert.Assert(direction >= 1 && direction <= 8, "direction should be in range of switch case")
	assert.Assert(x >= 0 && x < len(matrix), "x should be in range of matrix")
	assert.Assert(y >= 0 && y < len(matrix[0]), "y should be in range of matrix[x]")
	// 1. left -> right
	// 2. right -> left
	// 3. top -> bottom
	// 4. bottom -> top
	// 5. left-top -> right-bottom
	// 6. left-bottom -> right-top
	// 7. right-top -> left-bottom
	// 8. right-bottom -> left-top

	word := make([]byte, 4)

	switch direction {
	case 1:
		if x+3 >= len(matrix) {
			return 0
		}

		for i := 0; i < 4; i++ {
			word[i] = matrix[x+i][y]
		}
	case 2:
		if x-3 < 0 {
			return 0
		}

		for i := 0; i < 4; i++ {
			word[i] = matrix[x-i][y]
		}
	case 3:
		if y+3 >= len(matrix) {
			return 0
		}

		for i := 0; i < 4; i++ {
			word[i] = matrix[x][y+i]
		}
	case 4:
		if y-3 < 0 {
			return 0
		}

		for i := 0; i < 4; i++ {
			word[i] = matrix[x][y-i]
		}
	case 5:
		if y+3 >= len(matrix) || x+3 >= len(matrix) {
			return 0
		}

		for i := 0; i < 4; i++ {
			word[i] = matrix[x+i][y+i]
		}
	case 6:
		if y+3 >= len(matrix) || x-3 < 0 {
			return 0
		}

		for i := 0; i < 4; i++ {
			word[i] = matrix[x-i][y+i]
		}
	case 7:
		if y-3 < 0 || x+3 >= len(matrix) {
			return 0
		}

		for i := 0; i < 4; i++ {
			word[i] = matrix[x+i][y-i]
		}
	case 8:
		if y-3 < 0 || x-3 < 0 {
			return 0
		}

		for i := 0; i < 4; i++ {
			word[i] = matrix[x-i][y-i]
		}
	}

	return day4IsTargetWork(word)
}

func day4XmaxCount() {
	content := day4example2 // readInput("day4.txt")

	matrix := readMatrix(content)

	matrix.Print()

	// 1. left -> right
	// right -> left
	// top -> bottom
	// bottom -> top
	// left-top -> right-bottom
	// left-bottom -> right-top
	// right-top -> left-bottom
	// right-bottom -> left-top

	counter := 0

	for x := 0; x < len(matrix); x++ {
		line := matrix[x]
		for y := 0; y < len(line); y++ {
			for d := 1; d <= 8; d++ {
				counter += day4CheckCoordinate(x, y, matrix, d)
			}
		}
	}

	fmt.Println("Total:", counter, string(matrix[9][9]))
  matrix.Print()
}
