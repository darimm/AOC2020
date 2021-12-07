package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	aoc "github.com/darimm/AOCFunctions"
)

func sliceStrToFloat64(str []string) []float64 {
	var ret []float64
	for _, s := range str {
		i, _ := strconv.ParseFloat(s, 10)
		ret = append(ret, i)
	}
	return ret
}

func AbsFloat64(num float64) float64 {
	if num < 0 {
		return num * -1
	}
	return num
}

func findMaxFloat64(nums []float64) float64 {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func findMinFloat64(nums []float64) float64 {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func day7p1(data []string) {
	crabs := sliceStrToFloat64(strings.Split(data[0], ","))

	leftmostCrab := findMinFloat64(crabs)
	rightmostCrab := findMaxFloat64(crabs)
	fmt.Println(leftmostCrab, rightmostCrab)

	distanceSum := math.MaxFloat64
	distance := float64(0)
	for i := float64(0); i <= rightmostCrab; i++ {
		accumulator := float64(0)
		for _, v := range crabs {
			accumulator += math.Abs(v - i)
		}
		if accumulator < distanceSum {
			distanceSum = accumulator
			distance = i
		}
	}
	fmt.Printf("Crabs must move %f spaces, using %f fuel", distance, distanceSum)
}

func day7p2(data []string) {
	crabs := sliceStrToFloat64(strings.Split(data[0], ","))

	leftmostCrab := findMinFloat64(crabs)
	rightmostCrab := findMaxFloat64(crabs)
	fmt.Println(leftmostCrab, rightmostCrab)

	distanceSum := math.MaxFloat64
	distance := float64(0)
	for i := float64(0); i <= rightmostCrab; i++ {
		accumulator := float64(0)
		for _, v := range crabs {
			gauss := math.Abs(v - i)

			if gauss == 0 {
				continue
			}
			accumulator += (gauss / 2) * (1 + gauss)
		}
		if accumulator < distanceSum {
			distanceSum = accumulator
			distance = i
		}
	}
	fmt.Printf("Crabs must move %f spaces, using %f fuel", distance, distanceSum)
}

func main() {
	data, err := aoc.ReadFile(".\\input.txt")
	if err != nil {
		panic(err.Error)
	}

	day7p1(data)
	day7p2(data)
}
