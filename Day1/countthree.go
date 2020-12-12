package countthree

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func readFile(filePath string) (numbers []int, err error) {
	fd, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	var line int
	for {

		_, err = fmt.Fscanf(fd, "%d\n", &line)

		if err != nil {
			if err == io.EOF {
				err = nil
				return
			}
			fmt.Println(err)
			return nil, err
		}
		numbers = append(numbers, line)
	}
}

func countThree() {

	puzzleInput, err := readFile("input.txt")
	if err != nil {
		panic(fmt.Sprintf("Unable to read File - Error was: %+v\r\n", err))
	}

	sort.Sort(sort.IntSlice(puzzleInput))

	for k := range puzzleInput {
		i := puzzleInput[k]
		a := k + 1
		b := len(puzzleInput) - 1
		loopVar := true
		for loopVar {
			if a == b || a == len(puzzleInput) {
				loopVar = false
				continue
			}
			test := i + puzzleInput[a] + puzzleInput[b]
			if test < 2020 {
				a = a + 1
			}
			if test > 2020 {
				b = b - 1
			}
			if test == 2020 {
				fmt.Printf("%v\n", i*puzzleInput[a]*puzzleInput[b])
				loopVar = false
				break
			}
		}
	}
}
