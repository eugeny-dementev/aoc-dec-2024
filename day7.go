package main

import (
	"bufio"
	"fmt"
	"strings"
)

var day7example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func day7CalcPath() {
	input := readInput("day7.txt")

	var lines []string
	sc := bufio.NewScanner(strings.NewReader(string(input)))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	fmt.Println(lines)
}
