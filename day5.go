package main

import (
	"bytes"
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
	"strconv"
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

func day5ReadOrderingRules(rules []byte) map[int64][]int64 {
	lines := bytes.Split(rules, []byte("\n"))

	rulesMap := map[int64][]int64{}

	for _, line := range lines {
		values := bytes.Split(line, []byte("|"))
		left, err := strconv.ParseInt(string(values[0]), 0, 0)
		assert.NoError(err, "values[0] string should be parsed correctly to int", "values[0]", string(values[0]))

		right, err := strconv.ParseInt(string(values[1]), 0, 0)
		assert.NoError(err, "values[1] string should be parsed correctly to int", "values[1]", string(values[1]))

		rulesMap[left] = append(rulesMap[left], right)
	}

	return rulesMap
}

func day5ManualPrinting() {
	content := day5example // readInput("day5.txt")

	parts := bytes.Split(content, []byte("\n\n"))
	rules := day5ReadOrderingRules(parts[0])

	fmt.Printf("Content: %v\n", rules)
}
