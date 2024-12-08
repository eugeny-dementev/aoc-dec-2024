package main

import (
	"fmt"
)

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
	antenasSet       map[int]map[int]string
	antinodesSet     map[int]map[int]string
	antenasCoords    map[string][]Point
	board            [][]string
	antiNodesCounter int
}

func (m *Day8Map) getHeight() int { return len(m.board) }
func (m *Day8Map) getWidth() int  { return len(m.board[0]) }

func (m *Day8Map) getAntiNodeCount() int {
	counter := 0

	for _, in := range m.antinodesSet {
		for range in {
			counter++
		}
	}

	return counter
}

func (m *Day8Map) getSuperAntiNodeCounter() int {
  counter := 0

  for x, line := range m.board {
    for y := range line {
      if m.isSuperAntiNode(Point{x, y}) {
        counter++
      }
    }
  }

  return counter
}

func (m *Day8Map) addAntiNode(p Point) {
	if !m.isFree(p) || m.isOut(p) {
		return
	}

	_, ok := m.antinodesSet[p.x]
	if !ok {
		m.antinodesSet[p.x] = map[int]string{}
	}

	m.antinodesSet[p.x][p.y] = "#"

	m.antiNodesCounter++
}

func (m *Day8Map) isOut(p Point) bool {
	if p.x < 0 || p.x >= m.getHeight() {
		return true
	}

	if p.y < 0 || p.y >= m.getWidth() {
		return true
	}

	return false
}

/**
* Check for placed both antenas and antinodes
 */
func (m *Day8Map) isFree(p Point) bool {
	// innerAntenasMap, antenasOk := m.antenasSet[p.x]
	innerAntinodesMap, antinodesOk := m.antinodesSet[p.x]

	// if !antenasOk && !antinodesOk {
	if !antinodesOk {
		return true
	}

	// _, antenasOk = innerAntenasMap[p.y]
	_, antinodesOk = innerAntinodesMap[p.y]

	// return !antenasOk && !antinodesOk
	return !antinodesOk
}

func (m *Day8Map) isAntiNode(p Point) bool {
	_, ok := m.antinodesSet[p.x]

	if !ok {
		return false
	}

	_, ok = m.antinodesSet[p.x][p.y]

	return ok
}

func (m *Day8Map) isSuperAntiNode(p Point) bool {
  if m.isAntiNode(p) {
    return true
  }

  _, ok := m.antenasSet[p.x]

  if !ok {
    return false
  }

  _, ok = m.antenasSet[p.x][p.y]

  return ok
}

func (m *Day8Map) checkPair(p1 Point, points []Point) {
	if len(points) == 0 {
		return
	}

	for _, p2 := range points {
		xdiff := p1.x - p2.x
		ydiff := p1.y - p2.y

		a1x := p1.x + xdiff
		a1y := p1.y + ydiff

		a2x := p2.x - xdiff
		a2y := p2.y - ydiff

		for !m.isOut(Point{a1x, a1y}) {
			m.addAntiNode(Point{a1x, a1y})

			a1x += xdiff
			a1y += ydiff
		}

		for !m.isOut(Point{a2x, a2y}) {
			m.addAntiNode(Point{a2x, a2y})

			a2x -= xdiff
			a2y -= ydiff
		}
	}

	p1 = points[0]

	m.checkPair(p1, points[1:])
}

func (m *Day8Map) evaluate() {
	for _, coords := range m.antenasCoords {
		p1 := coords[0]
		m.checkPair(p1, coords[1:])
	}
}

func (m *Day8Map) printMap() {
	mapStr := ""
	for x, line := range m.board {
		for y, symbol := range line {
			if m.isAntiNode(Point{x, y}) {
				mapStr += "#"
			} else {
				mapStr += symbol
			}
		}
		mapStr += "\n"
	}

	fmt.Println(mapStr)
}

func day8CalcAntinodes() {
	// board := readMap(readLines(day8example))
	board := readFileMap("day8.txt")

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

	m := Day8Map{
		antenasSet:       antenasSet,
		antenasCoords:    antenasCoords,
		antinodesSet:     map[int]map[int]string{},
		board:            board,
		antiNodesCounter: 0,
	}

	// m.addAntiNode(Point{1, 7})

	// fmt.Println("antenasSet", antenasSet, m.isFree(Point{1, 7}))
	// fmt.Println("antenasCoords", antenasCoords)

	m.evaluate()

	m.printMap()
	fmt.Println("antinodes", m.getAntiNodeCount(), m.getSuperAntiNodeCounter())
}
