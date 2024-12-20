package main

import (
	"bufio"
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
	"strconv"
	"strings"
)

var day7example = []byte(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`)

type Context struct {
	equations    map[int64][]int64
	validAnswers map[int64]bool
}

func (c *Context) check(result, left, right int64, values []int64) {
	sumRes := left + right

	if sumRes == result && len(values) == 0 {
		c.validAnswers[result] = true
		return
	}

	mulRes := left * right

	if mulRes == result && len(values) == 0 {
		c.validAnswers[result] = true
		return
	}

	conResStr := strconv.FormatInt(left, 10) + strconv.FormatInt(right, 10)
	conRes, err := strconv.ParseInt(conResStr, 0, 0)
	assert.NoError(err, "should parse conResStr with no errors ", "conResStr", conResStr)

	if conRes == result && len(values) == 0 {
		c.validAnswers[result] = true
		return
	}

	if len(values) == 0 {
		return
	}

	right = values[0]

	c.check(result, sumRes, right, values[1:])
	c.check(result, mulRes, right, values[1:])
	c.check(result, conRes, right, values[1:])
}

func (c *Context) evaluate() int64 {
	for result, values := range c.equations {
		left := values[0]
		right := values[1]
		c.check(result, left, right, values[2:])
	}

	var result int64

	for value := range c.validAnswers {
		result += value
	}

	return result
}

func day7CalcPath() {
	input := readFileInput("day7.txt")

	context := Context{
		map[int64][]int64{},
		map[int64]bool{},
	}

	sc := bufio.NewScanner(strings.NewReader(string(input)))
	for sc.Scan() {
		parts := strings.Split(sc.Text(), ": ")
		assert.Assert(len(parts) == 2, "should have only two parts in single line", "parts", parts)

		result, err := strconv.ParseInt(parts[0], 0, 0)
		assert.NoError(err, "should parse parts[0] with no error", "parts[0]", parts[0])

		valuesStrings := strings.Split(parts[1], " ")
		var values []int64
		for _, valueStr := range valuesStrings {
			value, err := strconv.ParseInt(valueStr, 0, 0)
			assert.NoError(err, "should parse value with no error", "value", value)

			values = append(values, value)
		}

		context.equations[result] = values
	}

	result := context.evaluate()

	fmt.Println("Result:", result)
}
