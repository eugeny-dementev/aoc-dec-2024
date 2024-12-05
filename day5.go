package main

import (
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
	"math"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var day5example = []byte(`47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`)

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)
}

func day5ReadOrderingRules(rules string) map[string][]string {
	lines := strings.Split(rules, "\n")

	rulesMap := map[string][]string{}

	for _, line := range lines {
		if line[0] == byte(13) {
			continue
		}
		assert.Assert(len(line) > 1, "line should be none-0 length", "line str", string(line), "line byte", line)

		values := strings.Split(string(line), "|")
		assert.Assert(len(values) == 2, "there should be 2 values as result of split(|)", "values", values, "line", string(line))
		left := values[0]
		right := values[1]

		rulesMap[left] = append(rulesMap[left], right)
	}

	return rulesMap
}

var rulesMap map[string][]string

type Section []string

func (s Section) Len() int {
	return len(s)
}

func (s Section) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Section) Less(i, j int) bool {
	left := s[i]
	right := s[j]

	return slices.Contains(rulesMap[left], right)
}

func day5ManualPrinting() {
	content := readInput("day5.txt")

	parts := splitByEmptyNewline(string(content))
	rulesMap = day5ReadOrderingRules(parts[0])

	sections := strings.Split(parts[1], "\n")

	var middlesSum int64 = 0
	var incorrectMiddlesSum int64 = 0

	for _, section := range sections {
		pages := Section(strings.Split(section, ","))
		sort.Sort(pages)
		sortedSection := strings.Join(pages, ",")

		if section == sortedSection {
			length := len(pages)
			middle := int(math.Ceil(float64(length)/2)) - 1
			middleInt, err := strconv.ParseInt(pages[middle], 0, 0)
			assert.NoError(err, "should parse pages[middle] to int with no issues", "pages[middle]", pages[middle])

			// fmt.Println("Correct section:", section, sortedSection, middleInt, pages[middle])

			middlesSum += middleInt
		} else {
			values := strings.Split(sortedSection, ",")
			length := len(values)
			middle := int(math.Ceil(float64(length)/2)) - 1
			middleInt, err := strconv.ParseInt(values[middle], 0, 0)
			assert.NoError(err, "should parse values[middle] to int with no issues", "values[middle]", values[middle])

			// fmt.Println("Incorrect section:", section, sortedSection, middle, values[middle])

			incorrectMiddlesSum += middleInt
		}
	}

	fmt.Println("Sum of the middles:", middlesSum)
	fmt.Println("Sum of the incorrect middles:", incorrectMiddlesSum)
}
