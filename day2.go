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
		tolerable := day2IsReportTolerable(report, false)

		if safe {
			safeCounter++
		}

		if tolerable {
			tolerableCounter++
		}

		if !safe && tolerable {
			fmt.Println("Extra report:", report)
			day2IsReportTolerable(report, true)

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

func isIncreasing(report []int) (bool, bool) {
  // return true if most of the values are increasing
	prev := report[0]
	rest := report[1:]

  inc := 0

  for _, cur := range rest {
    if prev < cur {
      inc++
    } else {
      inc--
    }
  }

  if inc == 0 {
    return false, true
  }

  sign := inc / IntAbs(inc)

  return sign > 0, len(report) - IntAbs(inc) > 1
}

func day2IsReportTolerable(report []int, print bool) bool {
	prev := report[0]
	rest := report[1:]
  max := prev

	increasing, badOrder := isIncreasing(report)

  if badOrder {
    return false
  }

	orderFailsCounter := 0
	diffFailsCounter := 0
  maxFailsCounter := 0

	for _, cur := range rest {
		if increasing && prev > cur {
			orderFailsCounter += 1
		}

		if !increasing && prev < cur {
			orderFailsCounter += 1
		}

    if increasing && max > cur {
      maxFailsCounter += 1
    } else {
      max = cur
    }

    if !increasing && max < cur {
      maxFailsCounter += 1
    } else  {
      max = cur
    }


		diff := IntAbs(prev - cur)
		if diff < 1 || diff > 3 {
			diffFailsCounter += 1
		}

		prev = cur
	}

	if print {
		fmt.Println("Failing counters, (order, diff, max)", orderFailsCounter, diffFailsCounter, maxFailsCounter)
	}

	return orderFailsCounter+diffFailsCounter+maxFailsCounter < 2
}
