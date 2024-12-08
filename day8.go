package main

import "fmt"

var day8example = `............
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
............`

type Day8Map struct {
	antenasSet    map[int]map[int]string
	antenasCoords map[string][]Point
	board         [][]string
}

func day8CalcAntinodes() {
	board := readMap(readLines(day8example))
	// board := readFileMap("day8.txt")

	antenasSet := map[int]map[int]string{}
	antenasCoords := map[string][]Point{}

	for x, line := range board {
		for y, point := range line {
			if point != "." {
				_, ok := antenasSet[x]
				if !ok {
					antenasSet[x] = map[int]string{}
				}

				antenasSet[x][y] = point

				antenasCoords[point] = append(antenasCoords[point], Point{x, y})
			}
		}
	}

	fmt.Println("antenasSet", antenasSet)
	fmt.Println("antenasCoords", antenasCoords)
}
