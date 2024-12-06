package main

import (
	"fmt"
	"strings"
	"time"
)

var day6example = []byte(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)

type Point struct {
	x, y int
}

type Map struct {
	board [][]string
}

func (m *Map) getHeight() int {
	return len(m.board)
}

func (m *Map) getWidth() int {
	return len(m.board[0])
}

func (m *Map) getSymbol(x, y int) string {
	if x < 0 || y < 0 {
		return "."
	}

	if x >= m.getHeight() || y >= m.getWidth() {
		return "."
	}

	return m.board[x][y]
}

type Guard struct {
	place     *Point
	myMap     *Map
	direction string
	steps     int
  unique map[string]bool
}

func (g Guard) String() string {
	return fmt.Sprintf("Guard(at %v,%v facing %s)", g.place.x, g.place.y, g.direction)
}

func (g *Guard) isOut() bool {
	return g.place.x < 0 || g.place.x >= g.myMap.getHeight() || g.place.y < 0 || g.place.y >= g.myMap.getWidth()
}

func (g *Guard) step() {
  g.unique[fmt.Sprintf("%v:%v", g.place.x, g.place.y)] = true

	switch g.direction {
	case "^":
		g.place.x -= 1
	case ">":
		g.place.y += 1
	case "v":
		g.place.x += 1
	case "<":
		g.place.y -= 1
	}

	g.steps++
}

func (g *Guard) facingObstraction() bool {
	switch g.direction {
	case "^":
		return g.myMap.getSymbol(g.place.x-1, g.place.y) == "#"
	case ">":
		return g.myMap.getSymbol(g.place.x, g.place.y+1) == "#"
	case "v":
		return g.myMap.getSymbol(g.place.x+1, g.place.y) == "#"
	case "<":
		return g.myMap.getSymbol(g.place.x, g.place.y-1) == "#"
	}

	return false
}

func (g *Guard) turnRight() {
	switch g.direction {
	case "^":
		g.direction = ">"
	case ">":
		g.direction = "v"
	case "v":
		g.direction = "<"
	case "<":
		g.direction = "^"
	}
}

func (g *Guard) printMap() {
	for x, line := range g.myMap.board {
		for y, place := range line {
			if x == g.place.x && y == g.place.y {
				fmt.Print(g.direction)
			} else {
				fmt.Print(place)
			}
		}
		fmt.Print("\n")
	}
	time.Sleep(time.Millisecond * 100)
	for range g.myMap.board {
		fmt.Printf("\033[1A\033[K")
	}
}

func (g *Guard) startPatrol() {
	for !g.isOut() {
		//g.printMap()

		if g.facingObstraction() {
			g.turnRight()
		}

		g.step()
	}

	fmt.Println("Patrol is done, performed steps:", g.steps, len(g.unique))
}

func day6WalkAPath() {
	input := readInput("day6.txt")

	lines := strings.Split(string(input), "\r\n")

	myMap := &Map{[][]string{}}

	var guard Guard
	isGuardFound := false

	for x, line := range lines {
		places := strings.Split(line, "")
		myMap.board = append(myMap.board, places)

		if !isGuardFound {
			for y, place := range places {
				if place != "." && place != "#" {
					fmt.Println("Guard found", x, y)
					places[y] = "."
					guard = Guard{&Point{x, y}, myMap, place, 0, map[string]bool{}}

					isGuardFound = true
				}
			}
		}
	}

	guard.startPatrol()
}
