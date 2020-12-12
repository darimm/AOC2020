package xmascipher

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}

func readFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileBytes), "\n"), nil
}

func testNext(nums []int, index, preambleLength int) bool {
	testSlice := append([]int(nil), nums[index-preambleLength-1:index]...)
	lookupTable := make(map[int]int)

	for _, v := range testSlice {
		lookupTable[v] = nums[index] - v
	}

	for _, v := range lookupTable {
		if lookupTable[v] > 0 {
			return true
		}
	}
	return false
}

func findWeakness(nums []int, target int) (int, int) {
	for k := range nums {
		var total int
		low := 0
		high := 0
		for i := k; i < len(nums); i++ {
			total += nums[i]
			if nums[i] < low || low == 0 {
				low = nums[i]
			}
			if nums[i] > high {
				high = nums[i]
			}
			if total == target {
				return low, high
			}
			if total > target {
				break
			}
		}
	}
	return 0, 0
}

func main() {
	defer duration(track("main"))
	lines, err := readFile("input.txt")
	if err != nil {
		panic("Unable to read file")
	}

	var queue []int

	for _, v := range lines {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		queue = append(queue, i)
	}

	var testVal int
	for i := 26; i < len(queue); i++ {
		if !testNext(queue, i, 25) {
			fmt.Println(queue[i])
			testVal = queue[i]
			break
		}
	}
	min, max := findWeakness(queue, testVal)
	fmt.Println(min + max)
}
