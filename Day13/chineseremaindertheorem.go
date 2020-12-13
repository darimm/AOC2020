package chineseremaindertheorem

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func readFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileBytes), "\n"), nil
}

func AOCDay13p1(lines []string) {
	arrivalTime, _ := strconv.Atoi(lines[0])
	buses := strings.Split(lines[1], ",")
	minutesLate := make(map[int]int)

	for _, i := range buses {
		x, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		minutesLate[x] = -(arrivalTime % x) + x
	}

	lowest := math.MaxInt32
	bus := 0
	for k, v := range minutesLate {
		if v < lowest {
			lowest = v
			bus = k
		}
	}
	fmt.Println(lowest * bus)
}

func AOCDay13p2(lines []string) {
	buses := strings.Split(lines[1], ",")
	departureOffsets := make(map[int]int64)
	indices := []int{}

	for k, v := range buses {
		x, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		} else {
			indices = append(indices, k)
			departureOffsets[k] = x
		}

	}

	var myTime int64 = 0
	n1 := departureOffsets[0]
	for k, i := range indices {
		if k == 0 {
			continue
		}
		var multiplier int64
		for {
			multiplier++
			testTime := myTime + int64(multiplier)*n1
			test := (testTime + int64(i)) % departureOffsets[i]
			if test != 0 {
				continue
			}
			myTime = testTime
			n1 = n1 * departureOffsets[i]
			break
		}
	}
	fmt.Println(myTime)
}

func main() {
	lines, _ := readFile("input.txt")
	AOCDay13p1(lines)
	AOCDay13p2(lines)
}
