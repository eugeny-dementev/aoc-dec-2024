package main

import (
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
	"regexp"
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

func day5ManualPrinting() {
	content := day5example // readInput("day5.txt")

	parts := splitByEmptyNewline(string(content))
	fmt.Println(parts[0])
	rules := day5ReadOrderingRules(parts[0])

	fmt.Printf("Content: %v\n", rules)

	sections := strings.Split(parts[1], "\n")

	fmt.Println("Sections: ", sections)

	for _, section := range sections {
		pages := strings.Split(section, ",")

		fmt.Println("Pages:", pages)
	}
}
