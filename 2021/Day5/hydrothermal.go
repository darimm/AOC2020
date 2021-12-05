package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/darimm/AOCFunctions"
)

type coordinates struct {
	x, y int
}

type mapPoints struct {
	points map[coordinates]int
}

func (m *mapPoints) initialize() {
	m.points = make(map[coordinates]int)
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			m.points[coordinates{x, y}] = 0
		}
	}
}

func (m *mapPoints) addLineBasic(x1, y1, x2, y2 int) {
	if x1 == x2 {
		for y := aoc.MinInt(y1, y2); y <= aoc.MaxInt(y1, y2); y++ {
			m.points[coordinates{x1, y}]++
		}
	}
	if y1 == y2 {
		for x := aoc.MinInt(x1, x2); x <= aoc.MaxInt(x1, x2); x++ {
			m.points[coordinates{x, y1}]++
		}
	}
}

func (m *mapPoints) addLineAdvanced(x1, y1, x2, y2 int) {
	if x1 == x2 {
		for y := aoc.MinInt(y1, y2); y <= aoc.MaxInt(y1, y2); y++ {
			m.points[coordinates{x1, y}]++
		}
	}
	if y1 == y2 {
		for x := aoc.MinInt(x1, x2); x <= aoc.MaxInt(x1, x2); x++ {
			m.points[coordinates{x, y1}]++
		}
	}
	if x1 != x2 && y1 != y2 {
		var xmodifier, ymodifier int
		xmin := aoc.MinInt(x1, x2)
		xmax := aoc.MaxInt(x1, x2)

		if x1 > x2 {
			xmodifier = -1
		} else {
			xmodifier = 1
		}
		if y1 > y2 {
			ymodifier = -1
		} else {
			ymodifier = 1
		}
		for i := 0; i <= xmax-xmin; i++ {
			m.points[coordinates{x1 + (i * xmodifier), y1 + (i * ymodifier)}]++
		}
	}
}

func (m *mapPoints) overlappingPoints() int {
	total := 0
	for _, v := range m.points {
		if v > 1 {
			total++
		}
	}
	return total
}

func getVentCoordinates(data []string) [][]int {
	var ret [][]int

	for _, i := range data {
		if i == "" {
			continue
		}

		coordsRaw := strings.Split(i, " -> ")
		x1, _ := strconv.Atoi(strings.Split(coordsRaw[0], ",")[0])
		y1, _ := strconv.Atoi(strings.Split(coordsRaw[0], ",")[1])
		x2, _ := strconv.Atoi(strings.Split(coordsRaw[1], ",")[0])
		y2, _ := strconv.Atoi(strings.Split(coordsRaw[1], ",")[1])
		ret = append(ret, []int{x1, y1, x2, y2})
	}
	return ret
}

func day5p1(data []string) {
	var subMap mapPoints
	subMap.initialize()

	coords := getVentCoordinates(data)
	for _, i := range coords {
		subMap.addLineBasic(i[0], i[1], i[2], i[3])
	}

	fmt.Println("Number of Overlapping Horizontal and Vertical Lines: ", subMap.overlappingPoints())
}

func day5p2(data []string) {
	var subMap mapPoints
	subMap.initialize()

	coords := getVentCoordinates(data)
	for _, i := range coords {
		subMap.addLineAdvanced(i[0], i[1], i[2], i[3])
	}
	fmt.Println("Number of Overlapping Points: ", subMap.overlappingPoints())
}

func main() {
	data, err := aoc.ReadFile(".\\input.txt")
	if err != nil {
		panic(err.Error)
	}

	day5p1(data)
	day5p2(data)

}
