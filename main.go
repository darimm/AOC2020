package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	dead = iota
	alive
	floor
)

func readFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileBytes), "\n"), nil
}

func createBoard(lines []string) [][]int {
	var board [][]int

	for _, yval := range lines {
		if len(yval) == 0 {
			continue
		}
		var thisLine []int
		for _, xval := range yval {
			if xval == 'L' {
				thisLine = append(thisLine, dead)
			} else {
				thisLine = append(thisLine, floor)
			}
		}
		board = append(board, thisLine)
	}
	return board
}

func testLife(board [][]int, y, x int) int {
	xmin := 0
	xmax := len(board[0]) - 1
	ymin := 0
	ymax := len(board) - 1
	checkVal := board[y][x]
	emptyseats := 0
	fullseats := 0
	flr := 0
	for ytest := y - 1; ytest <= y+1; ytest++ {
		for xtest := x - 1; xtest <= x+1; xtest++ {
			fmt.Printf("y: %d, x: %d|", ytest, xtest)
			if ytest < ymin || ytest > ymax || xtest < xmin || xtest > xmax {
				fmt.Println("Out of Bounds")
				continue
			}
			if ytest == y && xtest == x {
				fmt.Println("Tested Cell")
				continue
			}
			fmt.Println(board[ytest][xtest])
			if (ytest == 7) && (xtest == 8) {
				fmt.Println("Cell value: ", board[ytest][xtest])
			}
			switch board[ytest][xtest] {
			case dead:
				emptyseats++
			case alive:
				fullseats++
			case floor:
				flr++
			}
		}
	}

	switch checkVal {
	case dead:
		if fullseats == 0 {
			fmt.Printf("Value of Cell Tested: Dead, returning Alive EmptySeats: %d FullSeats: %d Floor: %d\r\n", emptyseats, fullseats, flr)
			return alive
		}
		fmt.Printf("Value of Cell Tested: Dead, returning Dead EmptySeats: %d FullSeats: %d Floor: %d\r\n", emptyseats, fullseats, flr)
		return dead
	case alive:
		if fullseats >= 4 {
			fmt.Printf("Value of Cell Tested: Alive, returning Dead EmptySeats: %d FullSeats: %d Floor: %d\r\n", emptyseats, fullseats, flr)
			return dead
		}
		return alive
	case floor:
		fmt.Printf("Value of Cell Tested: Floor, returning Floor EmptySeats: %d FullSeats: %d Floor: %d\r\n", emptyseats, fullseats, flr)
		return floor
	}

	return -999 //this should never be reached
}

func applyRules11p1(board [][]int) [][]int {
	newboard := append([][]int(nil), board...)

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			fmt.Println("Testing: ", y, x)
			switch testLife(board, y, x) {
			case alive:
				newboard[y][x] = alive
			case dead:
				newboard[y][x] = dead
			case floor:
				newboard[y][x] = floor
			default:
				fmt.Println("-999")
			}
		}
	}
	return newboard
}

func countalive(board [][]int) int {
	returnVal := 0
	for _, yval := range board {
		for _, xval := range yval {
			if xval == alive {
				returnVal++
			}
		}
	}
	return returnVal
}

func printboard(board [][]int) {
	fmt.Println("    01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901")
	for y, yval := range board {
		fmt.Printf("%03d ", y)
		for _, xval := range yval {
			switch xval {
			case dead:
				fmt.Print("L")
			case alive:
				fmt.Print("#")
			case floor:
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	lines, err := readFile("input2.txt")
	if err != nil {
		panic("Unable to load file")
	}

	board := createBoard(lines)
	printboard(board)

	board = applyRules11p1(board)
	printboard(board)
	/*countLive := 0
	for {
		board = applyRules11p1(board)
		x := countalive(board)
		printboard(board)
		fmt.Println()
		if countLive == x {
			break
		}
		countLive = x
	}*/
}
