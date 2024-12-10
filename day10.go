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

  var day10example0 =`0123
1234
8765
9876`

var directionVectors map[rune]image.Point = map[rune]image.Point{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

var bounds image.Rectangle

var trailsCache map[image.Point][][]image.Point = map[image.Point][][]image.Point{}
var trailsNinesCache map[image.Point]map[image.Point]bool = map[image.Point]map[image.Point]bool{}

func walkTrail(cur image.Point, board [][]int, trail []image.Point, visitedSet map[image.Point]bool) {
	if !cur.In(bounds) {
		return
	}

	curHeight := board[cur.X][cur.Y]
	prev := trail[len(trail)-1]
	prevHeight := board[prev.X][prev.Y]

	if curHeight-1 != prevHeight {
		return
	}

	trail = append(trail, cur)

	if curHeight == 9 {
		visitedSet[cur] = true
		head := trail[0]
    trailsNinesCache[head][cur] = true
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

func day10PrintTrail(board [][]int, trail []image.Point) {
	trailSet := map[image.Point]bool{}
	for _, p := range trail {
		trailSet[p] = true
	}

	for x, line := range board {
		for y, height := range line {
			if (trailSet[image.Point{x, y}]) {
				fmt.Printf("%v", height)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func day10TrailsCount() {
	boardStrings := readFileMap("day10.txt")
	// boardStrings := readMap(readLines(day10example))
	// boardStrings := readMap(readLines(day10example0))

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
        trailsNinesCache[head] = map[image.Point]bool{}
			}

			line = append(line, int(height))
		}

		board = append(board, line)
	}

	bounds = image.Rectangle{image.Point{0, 0}, image.Point{len(board), len(board[0])}}

	for _, p := range startingPoints {
		walkTrail(p.Add(directionVectors['^']), board, []image.Point{p}, map[image.Point]bool{})
		walkTrail(p.Add(directionVectors['>']), board, []image.Point{p}, map[image.Point]bool{})
		walkTrail(p.Add(directionVectors['v']), board, []image.Point{p}, map[image.Point]bool{})
		walkTrail(p.Add(directionVectors['<']), board, []image.Point{p}, map[image.Point]bool{})
	}

	var score int
  var variationScore int

	for p, headMap := range trailsNinesCache {
		score += len(headMap)
    variationScore += len(trailsCache[p])
	}

	fmt.Println("Nines score", score)
  fmt.Println("Variation score", variationScore)
}
