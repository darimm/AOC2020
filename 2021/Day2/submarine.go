package main

import (
	"fmt"
	"io/ioutil"
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

func subDistance1(fileData []string) {
	subDepth := 0
	subDistance := 0
	for _, direction := range fileData {
		if direction == "" {
			continue
		}
		nav := strings.Split(direction, " ")
		t, _ := strconv.Atoi(nav[1])

		switch nav[0] {
		case "down":
			subDepth += t
		case "up":
			subDepth -= t
		case "forward":
			subDistance += t
		default:
			fmt.Printf("You shouldn't see this - nav[0] value is %s\r\n", nav[0])
		}
	}

	fmt.Printf("Final Depth: %d, Final Distance:%d\r\n", subDepth, subDistance)
	fmt.Printf("Result: %d\r\n", subDepth*subDistance)
}

func subDistance2(fileData []string) {
	subDepth := 0
	subDistance := 0
	subAim := 0
	for _, direction := range fileData {
		if direction == "" {
			continue
		}
		nav := strings.Split(direction, " ")
		t, _ := strconv.Atoi(nav[1])

		switch nav[0] {
		case "down":
			subAim += t
		case "up":
			subAim -= t
		case "forward":
			subDistance += t
			subDepth += (t * subAim)
		default:
			fmt.Printf("You shouldn't see this - nav[0] value is %s\r\n", nav[0])
		}
	}

	fmt.Printf("Final Depth: %d, Final Distance:%d\r\n", subDepth, subDistance)
	fmt.Printf("Result: %d\r\n", subDepth*subDistance)
}

func main() {
	fileData, err := readFile(".\\input.txt")
	if err != nil {
		panic(err.Error)
	}

	subDistance1(fileData)
	subDistance2(fileData)

}
