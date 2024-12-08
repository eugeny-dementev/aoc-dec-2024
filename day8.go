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
	antenasSet       map[int]map[int]string
  antinodesSet      map[int]map[int]string
	antenasCoords    map[string][]Point
	board            [][]string
	antiNodesCounter int
}

func (m *Day8Map) addAntiNode(p Point) {
	if !m.isFree(p) {
		return
	}

	_, ok := m.antinodesSet[p.x]
	if !ok {
		m.antinodesSet[p.x] = map[int]string{}
	}

  m.antinodesSet[p.x][p.y] = "#"
}

/**
* Check for placed both antenas and antinodes
 */
func (m *Day8Map) isFree(p Point) bool {
	innerAntenasMap, antenasOk := m.antenasSet[p.x]
  innerAntinodesMap, antinodesOk := m.antinodesSet[p.x]

	if !antenasOk && !antinodesOk {
		return true
	}

	_, antenasOk = innerAntenasMap[p.y]
	_, antinodesOk = innerAntinodesMap[p.y]

	return !antenasOk && !antinodesOk
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

	m := Day8Map{
		antenasSet:       antenasSet,
		antenasCoords:    antenasCoords,
    antinodesSet: map[int]map[int]string{},
		board:            board,
		antiNodesCounter: 0,
	}

  // m.addAntiNode(Point{1, 7})

	fmt.Println("antenasSet", antenasSet, m.isFree(Point{1, 7}))
	fmt.Println("antenasCoords", antenasCoords)
}
