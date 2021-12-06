package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/darimm/AOCFunctions"
)

func sliceStrToUInt64(str []string) []uint64 {
	var ret []uint64
	for _, s := range str {
		i, _ := strconv.ParseUint(s, 10, 64)
		ret = append(ret, i)
	}
	return ret
}

type lanternFish struct {
	school map[uint64]uint64
}

func (this *lanternFish) populate(fish []uint64) {
	this.school = make(map[uint64]uint64)
	for i := uint64(0); i < uint64(9); i++ {
		this.school[i] = 0
	}
	for _, v := range fish {
		this.school[v]++
	}
}

func (this *lanternFish) age() {
	newSchool := make(map[uint64]uint64)
	for i := uint64(0); i <= uint64(8); i++ {
		newSchool[i] = 0
	}
	for i := uint64(0); i <= uint64(8); i++ {
		if i == 0 {
			newSchool[uint64(8)] = this.school[uint64(0)]
			newSchool[uint64(6)] = this.school[uint64(0)]
		} else {
			newSchool[uint64(i-1)] += this.school[uint64(i)]
		}
	}
	this.school = newSchool
}

func main() {
	data, err := aoc.ReadFile(".\\input.txt")
	if err != nil {
		panic(err.Error)
	}

	fish := strings.Split(data[0], ",")
	fishAges := sliceStrToUInt64(fish)
	var lf lanternFish
	lf.populate(fishAges)

	for i := 1; i <= 80; i++ {
		lf.age()
	}
	finalFishNumbers := uint64(0)
	for _, v := range lf.school {
		finalFishNumbers += v
	}
	fmt.Println("Fish after 80 days: ", finalFishNumbers)

	for i := 1; i <= 176; i++ {
		lf.age()
	}
	finalFishNumbers = uint64(0)
	for _, v := range lf.school {
		finalFishNumbers += v
	}
	fmt.Println("Fish after 256 days: ", finalFishNumbers)
}
