package main

import "fmt"

var day8example = []byte(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`)

type Day8Map struct {
	antenas map[int]map[int]string
	board   [][]string
}

func day8CalcAntinodes() {
	board := readMap("day8.txt")

	antenas := map[int]map[int]string{}

	for x, line := range board {
		for y, point := range line {
			if point != "." {
        _, ok := antenas[x]
        if !ok {
          antenas[x] = map[int]string{}
        }

				antenas[x][y] = point
			}
		}
	}

  fmt.Println("antenas", antenas)
}
