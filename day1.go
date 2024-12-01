package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() ([]int, []int) {
	file, err := os.Open("day1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftArr := []int{}
	rightArr := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, " ")

		left := strings.TrimSpace(fields[0])
		right := strings.TrimSpace(fields[len(fields)-1])

		leftInt, err := strconv.ParseInt(left, 10, 64)
		if err != nil {
			panic(err)
		}
		leftArr = append(leftArr, int(leftInt))

		rightInt, err := strconv.ParseInt(right, 10, 64)
		if err != nil {
			panic(err)
		}
		rightArr = append(rightArr, int(rightInt))
	}

	return leftArr, rightArr
}

func day1part1distance() {
  left, right := readInput()

	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for i, v := range left {
		diff := v - right[i]
		if diff < 0 {
			sum -= diff
		} else {
			sum += diff
		}
	}

	fmt.Println("sum:", sum)
}

func day1part2similarityScore() {
  left, right := readInput()

	countMap := map[int]int{}
	for _, v := range right {
		countMap[v] += 1
	}

	similarityScore := 0

	for _, v := range left {
		score := v * countMap[v]

		similarityScore += score
	}

	fmt.Println("similarity score:", similarityScore)
}
