package main

import (
	"bufio"
	"fmt"
	"strconv"
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

func (m *Map) isObstacle(x, y int) bool {
	return m.getSymbol(x, y) == "#"
}

func (m *Map) cleanupMap() {
	cleaner := ""
	for range m.board {
		cleaner = cleaner + "\033[1A\033[K"
	}
	fmt.Print(cleaner)
}

func (m *Map) moveBackToRedraw() {
	mover := ""
	for range m.board {
		mover = mover + "\033[1A"
	}
	fmt.Print(mover)
}

func (m *Map) rotate() {
	matrix := m.board

	matrixSize := len(matrix)

	startRow := 0
	endRow := m.getWidth() - 1
	startColumn := 0
	endColumn := m.getHeight() - 1
	for x := 0; x < matrixSize; x++ {
		totalCycles := endRow - startRow

		for y := 0; y < totalCycles; y++ {
			temp := matrix[startRow][startColumn+y]
			matrix[startRow][startColumn+y] = matrix[endRow-y][startColumn]

			matrix[endRow-y][startColumn] = matrix[endRow][endColumn-y]

			matrix[endRow][endColumn-y] = matrix[startRow+y][endColumn]

			matrix[startRow+y][endColumn] = temp
		}

		startRow = startRow + 1
		endRow = endRow - 1
		startColumn = startColumn + 1
		endColumn = endColumn - 1
	}

	m.board = matrix
}

type Guard struct {
	place                   *Point
	myMap                   *Map
	unique                  map[string]bool
	direction               string
	lastThree               []Point
	possibleLoopObstacles   []Point
	confirmedLoopObstacles  []Point
	loopCompatibleObstacles []Point
	steps                   int
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

func (g *Guard) findLoopCompatibleObstaclesOnTheWay() {
	for _, obstaclePoint := range g.loopCompatibleObstacles {
		var newObstacleOnTheWay Point
		switch g.direction {
		case "v":
			if obstaclePoint.y < g.place.y && obstaclePoint.x < g.place.x {
				newObstacleOnTheWay = Point{g.place.x, obstaclePoint.y - 1}
			}
		case "<":
			if obstaclePoint.x < g.place.x && obstaclePoint.y > g.place.y {
				newObstacleOnTheWay = Point{obstaclePoint.x - 1, g.place.y}
			}
		case "^":
			if obstaclePoint.y > g.place.y && obstaclePoint.x > g.place.x {
				newObstacleOnTheWay = Point{g.place.x, obstaclePoint.y + 1}
			}
		case ">":
			if obstaclePoint.x > g.place.x && obstaclePoint.y < g.place.y {
				newObstacleOnTheWay = Point{obstaclePoint.x + 1, g.place.y}
			}
		}

		if !g.myMap.isObstacle(newObstacleOnTheWay.x, newObstacleOnTheWay.y) {
			g.possibleLoopObstacles = append(g.possibleLoopObstacles, newObstacleOnTheWay)
		}
	}
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

	if !g.myMap.isObstacle(possibleObstacle.x, possibleObstacle.y) {
		g.possibleLoopObstacles = append(g.possibleLoopObstacles, possibleObstacle)
	}

	g.findLoopCompatibleObstaclesOnTheWay()
}

func (g *Guard) confirmLoopObstacle(p Point) {
	g.confirmedLoopObstacles = append(g.confirmedLoopObstacles, p)
	g.loopCompatibleObstacles = append(g.loopCompatibleObstacles, g.lastThree...)
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

func (g *Guard) turnMap() { // for debugging
	g.myMap.rotate()

	curX, curY := g.place.x, g.place.y

	g.place.x = curY
	g.place.y = g.myMap.getWidth() - 1 - curX

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
	mapstring := ""
	for x, line := range g.myMap.board {
		for y, place := range line {
			if x == g.place.x && y == g.place.y {
				mapstring += g.direction
			} else if place == "#" && g.isOneOfLast3(x, y) {
				mapstring += "X"
			} else if g.isConfirmedLoopObstacle(x, y) {
				mapstring += "8"
			} else if g.isPossibleLoopObstacle(x, y) {
				mapstring += "O"
			} else {
				mapstring += place
			}
		}
		mapstring += "\n"
	}
	fmt.Print(mapstring)
}

func (g *Guard) printFrame() {
	mapstring := ""
	for x, line := range g.myMap.board {
		for y, place := range line {
			if x == g.place.x && y == g.place.y {
				mapstring += g.direction
			} else if place == "#" && g.isOneOfLast3(x, y) {
				mapstring += "X"
			} else if g.isConfirmedLoopObstacle(x, y) {
				mapstring += "8"
			} else if g.isPossibleLoopObstacle(x, y) {
				mapstring += "O"
			} else {
				mapstring += place
			}
		}
		mapstring += "\n"
	}
	fmt.Print(mapstring)
	time.Sleep(time.Millisecond * 500)
	if !g.isOut() {
		g.myMap.moveBackToRedraw()
	}
}

func (g *Guard) getUniqueConfirmedLoopObstacles() int {
	set := map[string]bool{}

	for _, obstaclePoint := range g.confirmedLoopObstacles {
		if obstaclePoint.x >= 0 && obstaclePoint.x < g.myMap.getHeight() && obstaclePoint.y >= 0 && obstaclePoint.y < g.myMap.getWidth() {
			xyStrs := []string{
				strconv.FormatInt(int64(obstaclePoint.x), 10),
				strconv.FormatInt(int64(obstaclePoint.y), 10),
			}
			key := strings.Join(xyStrs, ":")
			set[key] = true
		}
	}

	return len(set)
}

func (g *Guard) startPatrol(visualize bool) {
	for !g.isOut() {
		if visualize {
			g.printFrame()
		}

		if g.facingObstraction() {
			g.turnRight()
		} else {
			g.step()

			if g.isPossibleLoopObstacle(g.place.x, g.place.y) {
				g.confirmLoopObstacle(Point{g.place.x, g.place.y})
			}
		}
	}

	fmt.Println("Patrol is done, performed steps:", g.steps, len(g.unique))
	fmt.Println("Confirmed loop obstacles found:", g.getUniqueConfirmedLoopObstacles())

	g.printMap()
}

func day6WalkAPath() {
	input := readInput("day6.txt")

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
						place:                   &Point{x, y},
						myMap:                   myMap,
						direction:               place,
						steps:                   0,
						unique:                  map[string]bool{},
						lastThree:               []Point{},
						possibleLoopObstacles:   []Point{},
						confirmedLoopObstacles:  []Point{},
						loopCompatibleObstacles: []Point{},
					}

					isGuardFound = true
				}
			}
		}
	}

	guard.printMap()
	guard.startPatrol(true)
}
