package main

import "fmt"

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
