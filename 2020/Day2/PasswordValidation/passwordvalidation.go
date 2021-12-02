package passwordvalidation

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type passwordLine struct {
	min          int
	max          int
	chartoAppear string
	pw           string
}

func readFile(filePath string) (lines []passwordLine, err error) {
	fd, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	var min, max int
	var charToAppear string
	var password string

	var charRange string

	for {

		_, err = fmt.Fscanf(fd, "%s %s %s\n", &charRange, &charToAppear, &password)

		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println(err)
			return nil, err
		}

		tempSlice := strings.Split(charRange, "-")
		min, _ = strconv.Atoi(tempSlice[0])
		max, _ = strconv.Atoi(tempSlice[1])
		charToAppear = strings.Replace(charToAppear, ":", "", -1)

		lines = append(lines, passwordLine{
			min:          min,
			max:          max,
			chartoAppear: charToAppear,
			pw:           password,
		})
	}
}

func main() {
	data, err := readFile("input.txt")
	if err != nil && err != io.EOF {
		panic(fmt.Sprintf("Error: %+v", err))
	}
	successCount := 0
	for _, entry := range data {
		if strings.Count(entry.pw, entry.chartoAppear) < entry.min || strings.Count(entry.pw, entry.chartoAppear) > entry.max {
			continue
		}
		successCount++
	}
	fmt.Println(successCount)
}
