package main

import (
	"bufio"
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readReportsInput() [][]int {
	file, err := os.Open("day2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var reports [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		strValues := strings.Split(line, " ")

		report := []int{}

		for _, strValue := range strValues {
			intValue, err := strconv.ParseInt(strValue, 0, 0)
			assert.NoError(err, "strings in day2.txt file should all be parsable integers", "strValue", strValue)

			report = append(report, int(intValue))
		}

		reports = append(reports, report)
	}

	return reports
}

func day2CountSafeReports() {
	reports := readReportsInput()

	safeCounter := 0
	tolerableCounter := 0

	for _, report := range reports {
		safe := day2IsReportSafe(report)
		tolerable := day2IsReportTolerable(report)

		if safe {
			safeCounter++
		}

		if tolerable {
			tolerableCounter++
		}

		if !safe && tolerable {
			day2IsReportTolerable(report)
		}
	}

	fmt.Println("Total safe reports:", safeCounter, "out of", len(reports))
	fmt.Println("Total tolerable reports:", tolerableCounter, "out of", len(reports))
}

func IntAbs(x int) int {
	if x < 0 {
		return -1 * x
	}

	return x
}

func day2IsReportSafe(report []int) bool {
	prev := report[0]
	last := report[len(report)-1]
	rest := report[1:]

	increasing := prev < last

	for _, cur := range rest {
		if increasing && prev > cur {
			return false
		}

		if !increasing && prev < cur {
			return false
		}

		diff := IntAbs(prev - cur)
		if diff < 1 || diff > 3 {
			return false
		}

		prev = cur
	}

	return true
}

func day2IsReportTolerable(report []int) bool {
	for i := range report {
		clone := make([]int, len(report))
		copy(clone, report)
		reduced := append(clone[:i], clone[i+1:]...)

		if day2IsReportSafe(reduced) {
			return true
		}
	}

	return false
}
