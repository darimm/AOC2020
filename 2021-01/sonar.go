package sonar

import (
	"fmt"
	"io"
	"os"
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

func main() {
	input, err := readFile(".\\input.txt")
	if err != nil {
		panic("Something went wrong")
	}

	increases := 0

	for index, value := range input {
		if index == 0 {
			continue
		}
		if value > input[index-1] {
			increases++
		}
	}
	fmt.Println(increases)

	increaseWindow := 0
	for index, value := range input {
		if index == 0 || index == 1 || index == 2 {
			continue
		}
		currentWindow := value + input[index-1] + input[index-2]
		previousWindow := input[index-1] + input[index-2] + input[index-3]

		if currentWindow > previousWindow {
			increaseWindow++
		}
	}
	fmt.Println(increaseWindow)
}
