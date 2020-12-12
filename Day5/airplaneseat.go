package airplaneseat

import (
	"fmt"
	"io/ioutil"
	"sort"
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

func lower(low, high int) (int, int) {
	median := (high - low + 1) / 2
	return low, high - median
}

func higher(low, high int) (int, int) {
	median := (high - low + 1) / 2
	return low + median, high
}

func getSeatID(bsp string) (int, error) {
	if len(bsp) != 10 {
		return 0, fmt.Errorf("Invalid String")
	}
	row, err := strconv.ParseInt(strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1").Replace(bsp), 2, 16)
	if err != nil {
		return 0, err
	}
	return int(row), nil
}

func getSeatID3(bsp string) (int, error) {
	if len(bsp) != 10 {
		return 0, fmt.Errorf("Invalid String")
	}
	row, err := strconv.ParseInt(strings.Replace(strings.Replace(bsp[0:7], "F", "0", -1), "B", "1", -1), 2, 8)
	if err != nil {
		return 0, err
	}
	col, err := strconv.ParseInt(strings.Replace(strings.Replace(bsp[7:], "L", "0", -1), "R", "1", -1), 2, 8)
	if err != nil {
		return 0, err
	}
	return int(row*8 + col), nil
}

func getSeatID2(bsp string) (int, error) {
	if len(bsp) != 10 {
		return 0, fmt.Errorf("Invalid String")
	}
	startMax := 127
	startMin := 0
	rowMin := 0
	rowMax := 7

	for _, x := range bsp {
		switch x {
		case 'F':
			startMin, startMax = lower(startMin, startMax)
		case 'B':
			startMin, startMax = higher(startMin, startMax)
		case 'L':
			rowMin, rowMax = lower(rowMin, rowMax)
		case 'R':
			rowMin, rowMax = higher(rowMin, rowMax)
		}
	}
	return (startMax * 8) + rowMax, nil
}

func main() {
	data, err := readFile("input.txt")
	if err != nil {
		panic("Unable to read file")
	}
	highest := 0
	var seats []int
	for _, v := range data {
		seat, err := getSeatID(v)
		if err != nil {
			continue
		}
		seats = append(seats, seat)
		if seat > highest {
			highest = seat
		}
	}
	println("Highest: ", highest)

	sort.Sort(sort.IntSlice(seats))
	lastSeatID := seats[0] - 1
	for _, x := range seats {
		if x-1 != lastSeatID {
			println("My Seat ID:", x-1)
			break
		} else {
			lastSeatID = x
		}
	}
}
