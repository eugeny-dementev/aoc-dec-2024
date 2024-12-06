package main

import (
	"bufio"
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
	place                  *Point
	myMap                  *Map
	unique                 map[string]bool
	direction              string
	lastThree              []Point
	possibleLoopObstacles  []Point
	confirmedLoopObstacles []Point
	steps                  int
}

func (g Guard) String() string {
	return fmt.Sprintf("Guard(at %v,%v facing %s)", g.place.x, g.place.y, g.direction)
}

func (g *Guard) isOneOfLast3(x, y int) bool {
	for _, obstaclePoint := range g.lastThree {
		if obstaclePoint.x == x && obstaclePoint.y == y {
			return true
		}
	}

	return false
}

func (g *Guard) isPossibleLoopObstacle(x, y int) bool {
	for _, obstaclePoint := range g.possibleLoopObstacles {
		if obstaclePoint.x == x && obstaclePoint.y == y {
			return true
		}
	}

	return false
}

func (g *Guard) isConfirmedLoopObstacle(x, y int) bool {
	for _, obstaclePoint := range g.confirmedLoopObstacles {
		if obstaclePoint.x == x && obstaclePoint.y == y {
			return true
		}
	}

	return false
}

func (g *Guard) registerLastObstacle(p Point) {
	if len(g.lastThree) < 3 {
		g.lastThree = append(g.lastThree, p)
	} else {
		g.lastThree = append(g.lastThree[1:], p)
	}
}

func (g *Guard) calculatePossibleLoopObstacle() {
	if len(g.lastThree) < 3 {
		return
	}

	firstOfThree := g.lastThree[0]
	lastOfThree := g.lastThree[2]

	possibleObstacle := Point{}

	switch g.direction {
	case "^":
		possibleObstacle.x = lastOfThree.x + 1
		possibleObstacle.y = firstOfThree.y + 1
	case ">":
		possibleObstacle.x = firstOfThree.x + 1
		possibleObstacle.y = lastOfThree.y - 1
	case "v":
		possibleObstacle.x = lastOfThree.x - 1
		possibleObstacle.y = firstOfThree.y - 1
	case "<":
		possibleObstacle.x = firstOfThree.x - 1
		possibleObstacle.y = lastOfThree.y + 1
	}

	g.possibleLoopObstacles = append(g.possibleLoopObstacles, possibleObstacle)
}

func (g *Guard) confirmLoopObstacle(p Point) {
	g.confirmedLoopObstacles = append(g.confirmedLoopObstacles, p)
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
	xShift, yShift := g.place.x, g.place.y
	switch g.direction {
	case "^":
		xShift -= 1
	case ">":
		yShift += 1
	case "v":
		xShift += 1
	case "<":
		yShift -= 1
	}

	symbol := g.myMap.getSymbol(xShift, yShift)

	if symbol == "#" {
		g.registerLastObstacle(Point{xShift, yShift})

		if len(g.lastThree) == 3 {
			g.calculatePossibleLoopObstacle()
		}

		return true
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
			} else if place == "#" && g.isOneOfLast3(x, y) {
				fmt.Print("X")
			} else if g.isConfirmedLoopObstacle(x, y) {
				fmt.Print("8")
			} else if g.isPossibleLoopObstacle(x, y) {
				fmt.Print("O")
			} else {
				fmt.Print(place)
			}
		}
		fmt.Print("\n")
	}
	time.Sleep(time.Millisecond * 100)
	if !g.isOut() {
		for range g.myMap.board {
			fmt.Printf("\033[1A\033[K")
		}
	}
}

func (g *Guard) startPatrol() {
	for !g.isOut() {
		g.printMap()

		if g.facingObstraction() {
			g.turnRight()
		}

		g.step()

		if g.isPossibleLoopObstacle(g.place.x, g.place.y) {
			g.confirmLoopObstacle(Point{g.place.x, g.place.y})
		}
	}

	fmt.Println("Patrol is done, performed steps:", g.steps, len(g.unique))
	fmt.Println("Confirmed loop obstacles found:", len(g.confirmedLoopObstacles))

	g.printMap()
}

func day6WalkAPath() {
	input := day6example // readInput("day6.txt")

	var lines []string
	sc := bufio.NewScanner(strings.NewReader(string(input)))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

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
					guard = Guard{
						place:                  &Point{x, y},
						myMap:                  myMap,
						direction:              place,
						steps:                  0,
						unique:                 map[string]bool{},
						lastThree:              []Point{},
						possibleLoopObstacles:  []Point{},
						confirmedLoopObstacles: []Point{},
					}

					isGuardFound = true
				}
			}
		}
	}

	guard.startPatrol()
}
