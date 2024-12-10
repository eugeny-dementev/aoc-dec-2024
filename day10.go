package main

import (
	assert "eugeny-dementev/aoc-dec-2024/pkg"
	"fmt"
	"image"
	"strconv"
)

var day10example = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

var directionVectors map[rune]image.Point = map[rune]image.Point{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

var bounds image.Rectangle

var trailsCache map[image.Point][][]image.Point = map[image.Point][][]image.Point{}

func walkTrail(cur image.Point, board [][]int, trail []image.Point, visitedSet map[image.Point]bool) {
	if !cur.In(bounds) || visitedSet[cur] {
		return
	}

	curHeight := board[cur.X][cur.Y]
	prev := trail[len(trail)-1]
	prevHeight := board[prev.X][prev.Y]

	visitedSet[cur] = true

	if curHeight-1 != prevHeight {
		return
	}

	trail = append(trail, cur)

	if curHeight == 9 {
		head := trail[0]
		trailsCache[head] = append(trailsCache[head], trail)

		return
	}

	walkTrail(cur.Add(directionVectors['^']), board, trail, visitedSet)
	walkTrail(cur.Add(directionVectors['>']), board, trail, visitedSet)
	walkTrail(cur.Add(directionVectors['v']), board, trail, visitedSet)
	walkTrail(cur.Add(directionVectors['<']), board, trail, visitedSet)
}

func day10PrintMap(board [][]int) {
	for _, line := range board {
		for _, height := range line {
			fmt.Printf("%v", height)
		}
		fmt.Printf("\n")
	}
}

func day10TrailsCount() {
	// board := readFileMap("day10.txt")
	boardStrings := readMap(readLines(day10example))

	board := [][]int{}

	startingPoints := []image.Point{}

	for x, lineString := range boardStrings {
		line := []int{}
		for y, place := range lineString {
			height, err := strconv.ParseInt(place, 0, 0)
			assert.NoError(err, "should parse height from place with no error", "place", place)

			if height == 0 {
				head := image.Point{x, y}
				startingPoints = append(startingPoints, head)
				trailsCache[head] = [][]image.Point{}
			}

			line = append(line, int(height))
		}

		board = append(board, line)
	}

	bounds = image.Rectangle{image.Point{0, 0}, image.Point{len(board), len(board[0])}}

  day10PrintMap(board)

	fmt.Println(bounds)
	fmt.Println(startingPoints)

	for _, p := range startingPoints {
		walkTrail(p.Add(directionVectors['^']), board, []image.Point{p}, map[image.Point]bool{})
		walkTrail(p.Add(directionVectors['>']), board, []image.Point{p}, map[image.Point]bool{})
		walkTrail(p.Add(directionVectors['v']), board, []image.Point{p}, map[image.Point]bool{})
		walkTrail(p.Add(directionVectors['<']), board, []image.Point{p}, map[image.Point]bool{})
	}

	fmt.Println(trailsCache)
}
