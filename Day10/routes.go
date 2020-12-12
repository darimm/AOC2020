package routes

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type graph map[int][]int

func readFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileBytes), "\n"), nil
}

func convertStringSliceToInt(s []string) []int {
	var returnVal []int
	for _, i := range s {
		v, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		returnVal = append(returnVal, v)
	}
	return returnVal
}

func convertIntSliceToGraph(i []int) graph {
	returnVal := make(graph)
	for k, v := range i {
		s := []int{}
		for _, j := range i[k+1:] {
			if j-v <= 3 {
				s = append(s, j)
			} else {
				returnVal[v] = append([]int(nil), s...)
				break
			}
		}
		if returnVal[v] == nil {
			returnVal[v] = append([]int(nil), s...)
		}
	}
	return returnVal
}

func rCountPaths(start, destination, pathCount int, g graph) int {
	if start == destination {
		pathCount++
	} else {
		if g[start] != nil {
			for _, i := range g[start] {
				pathCount = rCountPaths(i, destination, pathCount, g)
			}
		}
	}
	return pathCount
}

//someone else's solution
func countPossibleAdapters(fromIndex int, nums []int, visited map[int]int) int {
	if fromIndex >= len(nums)-3 {
		return 1
	}

	num := nums[fromIndex]
	if res, ok := visited[num]; ok {
		return res
	}

	var count int
	for i := fromIndex + 1; i < fromIndex+4; i++ {
		n := nums[i]
		if areCompatible(num, n) {
			count += countPossibleAdapters(i, nums, visited)
		}
	}

	visited[num] = count // store the result
	return count
}

func areCompatible(low, high int) bool {
	return low+1 == high || low+2 == high || low+3 == high
}

//end someone else's solution

func main() {
	data, err := readFile("input.txt")
	if err != nil {
		panic("Unable to load file")
	}

	adapters := convertStringSliceToInt(data)
	adapters = append(adapters, 0)
	sort.Sort(sort.IntSlice(adapters))
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	fmt.Println(adapters)
	one := 0
	two := 0
	three := 0
	for i := range adapters {
		if i == 0 {
			continue
		}
		switch adapters[i] - adapters[i-1] {
		case 1:
			one++
		case 2:
			two++
		case 3:
			three++
		}
	}
	fmt.Println(one * three)

	//myGraph := convertIntSliceToGraph(adapters)

	temp := make(map[int]int)
	n := countPossibleAdapters(0, adapters, temp)
	fmt.Println(temp)
	//n := rCountPaths(0, adapters[len(adapters)-1], 0, myGraph)

	fmt.Println(n)
	//fmt.Println(myGraph)

}
