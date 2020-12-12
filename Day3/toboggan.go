package toboggan

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileBytes), "\n"), nil
}

func tobogganRide(mapData []string, x, y int) (returnVal int) {
	//fmt.Println("0123456789012345678901234567890")
	maxY := len(mapData) - 1
	returnVal = 0
	xPos := x
	for n := y; n < maxY; n = n + y {
		//there's a blank line at the bottom of the input file
		if len(mapData[n]) == 0 {
			break
		}
		if mapData[n][xPos] == '#' {
			returnVal++
		}
		//fmt.Printf("%s MapChar: %s StrPos: %d Tree: %t Row: %d\r\n", mapData[n], string(mapData[n][xPos]), xPos, mapData[n][xPos] == '#', n)
		xPos += x
		if xPos >= len(mapData[n-1]) {
			xPos = xPos - len(mapData[n-1])
		}
	}
	return
}

func main() {
	myMap, err := readFile("input.txt")
	if err != nil {
		panic("Unable to load map data")
	}

	result := tobogganRide(myMap, 3, 1)
	fmt.Println(result)
	result *= tobogganRide(myMap, 1, 1)
	fmt.Println(result)
	result *= tobogganRide(myMap, 5, 1)
	fmt.Println(result)
	result *= tobogganRide(myMap, 7, 1)
	fmt.Println(result)
	result *= tobogganRide(myMap, 1, 2)
	fmt.Println(result)
}
