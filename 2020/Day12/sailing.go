package sailing

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	directionNorth = 0
	directionSouth = 180
	directionEast  = 90
	directionWest  = 270
)

type instruction struct {
	cmd    rune
	number int
}

func readFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileBytes), "\n"), nil
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func parsedirection(directions []string) map[int]instruction {
	returnVal := make(map[int]instruction)
	for k, v := range directions {
		if len(v) == 0 {
			continue
		}
		x, _ := strconv.Atoi(v[1:])
		returnVal[k] = instruction{
			cmd:    rune(v[0]),
			number: x,
		}
	}
	return returnVal
}

func aoc2020d12p1(lines []string) {
	directions := parsedirection(lines)
	ns := 0
	ew := 0
	facing := directionEast
	for i := 0; i <= len(directions); i++ {
		switch directions[i].cmd {
		case 'N':
			ns += directions[i].number
		case 'S':
			ns -= directions[i].number
		case 'E':
			ew += directions[i].number
		case 'W':
			ew -= directions[i].number
		case 'R':
			facing = (facing + directions[i].number) % 360
		case 'L':
			facing = facing - directions[i].number
			if facing < 0 {
				facing += 360
			}
		case 'F':
			switch facing {
			case directionNorth:
				ns += directions[i].number
			case directionSouth:
				ns -= directions[i].number
			case directionEast:
				ew += directions[i].number
			case directionWest:
				ew -= directions[i].number
			}
		}
	}
	result := absInt(ns) + absInt(ew)
	fmt.Println(result)
}

type waypoint struct {
	ew int
	ns int
}

func aoc2020d12p2(lines []string) {
	directions := parsedirection(lines)
	wp := waypoint{ew: 10, ns: 1}
	ns := 0
	ew := 0
	for i := 0; i <= len(directions); i++ {
		switch directions[i].cmd {
		case 'N':
			wp.ns += directions[i].number
		case 'S':
			wp.ns -= directions[i].number
		case 'E':
			wp.ew += directions[i].number
		case 'W':
			wp.ew -= directions[i].number
		case 'R':
			switch directions[i].number {
			case 90:
				wp.ew, wp.ns = wp.ns, -wp.ew
			case 180:
				wp.ew, wp.ns = -wp.ew, -wp.ns
			case 270:
				wp.ew, wp.ns = -wp.ns, wp.ew
			}
		case 'L':
			switch directions[i].number {
			case 90:
				wp.ew, wp.ns = -wp.ns, wp.ew
			case 180:
				wp.ew, wp.ns = -wp.ew, -wp.ns
			case 270:
				wp.ew, wp.ns = wp.ns, -wp.ew
			}
		case 'F':
			ns += directions[i].number * wp.ns
			ew += directions[i].number * wp.ew
		}
	}
	result := absInt(ns) + absInt(ew)
	fmt.Println(result)

}

func main() {
	lines, err := readFile("input.txt")
	if err != nil {
		panic("Unable to load input.txt")
	}

	aoc2020d12p1(lines)
	aoc2020d12p2(lines)
}
