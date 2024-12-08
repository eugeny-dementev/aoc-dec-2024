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

var dot = []byte(".")[0]

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

	if isTarget {
		return 1
	}

	return 0
}

func day4CheckCoordinate(x, y, direction int, matrix Matrix) int {
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

func day4CheckXMaxForCoordinate(x, y int, matrix Matrix) int {
	assert.Assert(x >= 1 && x < len(matrix)-1, "x should be in range of matrix -1 on both edges")
	assert.Assert(y >= 1 && y < len(matrix[0])-1, "y should be in range of matrix[x] -1 on both edges")

	lineLen := len(matrix[0])
	for i, line := range matrix {
		curLineLen := len(line)
		assert.Assert(curLineLen == lineLen, "each line should have same length as previous", "curLineLen", curLineLen, "i", i, "len(matrix)", len(matrix))
		lineLen = curLineLen
	}

	// 1. left -> right
	// 2. right <- left
	// 3. top -> bottom
	// 4. bottom -> top

	ltrb, lbrt := string([]byte{matrix[x-1][y-1], A, matrix[x+1][y+1]}), string([]byte{matrix[x+1][y-1], A, matrix[x-1][y+1]})

	matched := 0

	if ltrb == "MAS" && lbrt == "MAS" {
		matched = 1
	}

	if ltrb == "MAS" && lbrt == "SAM" {
		matched = 1
	}

	if ltrb == "SAM" && lbrt == "MAS" {
		matched = 1
	}

	if ltrb == "SAM" && lbrt == "SAM" {
		matched = 1
	}

	// subMatrix := Matrix{
	// 	[]byte{matrix[x-1][y-1], dot, matrix[x-1][y+1]},
	// 	[]byte{dot, A, dot},
	// 	[]byte{matrix[x+1][y-1], dot, matrix[x+1][y+1]},
	// }

	// if matched == 1 {
	// 	subMatrix.Print()
	// }

	return matched
}

func day4XmaxCount() {
	content := readFileInput("day4.txt")

	matrix := readMatrix(content)

	XmasCounter := 0
	XMasCounter := 0

	for x := 0; x < len(matrix); x++ {
		line := matrix[x]
		for y := 0; y < len(line); y++ {
			if line[y] == X {
				for d := 1; d <= 8; d++ {
					XmasCounter += day4CheckCoordinate(x, y, d, matrix)
				}
			}
		}
	}

	for x := 1; x < len(matrix)-1; x++ {
		for y := 1; y < len(matrix[x])-1; y++ {
			if matrix[x][y] == A {
				XMasCounter += day4CheckXMaxForCoordinate(x, y, matrix)
			}
		}
	}

	fmt.Println("Total:", XmasCounter)
	fmt.Println("Total:", XMasCounter)
}
