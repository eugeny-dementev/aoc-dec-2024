package main

import (
	"fmt"
	"strconv"
)

var day9example = `2333133121414131402`

func day9TwoPointer() {
	input := string(readFileInput("day9.txt"))
	// input := day9example

	var data []int

	id := 0
	dot := -1

	for i, c := range input {
		num := int(c) - '0'

		if i%2 == 0 {
			for range num {
				data = append(data, id)
			}
			id++
		} else {
			for range num {
				data = append(data, dot)
			}
		}
	}

	// fmt.Println("data:", data)

	left := 0

	right := len(data) - 1

	for left < right {

		for data[left] != dot {
			left++
		}

		for data[right] == dot {
			right--
		}

		if left >= right {
			break
		}

		data[left], data[right] = data[right], data[left]

		left++
		right--
	}

	left = 0

	var i, hashsum int64

	for data[i] != dot {
		c := data[i]

		hashsum += i * int64(c)

		i++
	}

	fmt.Println("Hashsum", hashsum)
}

func printLine(data []Day9Section) string {
	var str string

	for _, section := range data {
		for range section.size {
			if section.sectionType == DOT {
				str += "."
			} else {
				str += strconv.FormatInt(int64(section.id), 10)
			}
		}
	}

	return str
}

var DOT, DATA = -1, 1

type Day9Section struct {
	sectionType int // -1 if "." and 1 if data
	size        int
	id          int // -1 if DOT
}

func (s Day9Section) String() string {
	t := "."
	if s.sectionType == DATA {
		t = strconv.FormatInt(int64(s.id), 10)
	}

	return fmt.Sprintf("(%v|%v)", t, s.size)
}

func day9Optimised() {
	// input := string(readFileInput("day9.txt"))
	input := day9example

	var data []Day9Section

	id := 0

	for i, c := range input {
		size := int(c) - '0'

		if i%2 == 0 {
			data = append(data, Day9Section{DATA, size, id})
			id++
		} else if size > 0 {
			data = append(data, Day9Section{DOT, size, DOT})
		}
	}

	fmt.Println("data:", printLine(data))

	left := 0

	right := len(data) - 1

	protector := len(data)

	for right >= 0 && protector >= 0 {
		for data[right].sectionType == DOT {
			right--
		}

		fmt.Println("right is now", data[right])

    leftLoop := true
		for leftLoop || (data[left].sectionType == DATA || data[right].size > data[left].size) {
			fmt.Println("left++", left)
			left++

      if left > len(data)-1 {
        right--
        left = 0
        leftLoop = false
      }
		}

    if leftLoop {
      protector++
      continue
    }

		fmt.Println("left is now", data[left])

		if left >= right {
			break
		}

		diff := data[left].size - data[right].size

		fmt.Println("data before swap", printLine(data))
		data[left] = data[right]
		data[right] = Day9Section{DOT, data[left].size, DOT}
		if diff > 0 {
			pos := left + 1
			data = append(data[:pos], append([]Day9Section{{DOT, diff, DOT}}, data[pos:]...)...)
		}
		fmt.Println("data after swap", printLine(data))

		left = 0

		protector--
	}

	fmt.Println("data:", printLine(data))

	// left = 0

	// var i, hashsum int64

	// for data[i] != dot {
	// 	c := data[i]

	// 	hashsum += i * int64(c)

	// 	i++
	// }

	// fmt.Println("Hashsum", hashsum)
}
