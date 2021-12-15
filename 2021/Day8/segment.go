package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	aoc "github.com/darimm/AOCFunctions"
)

func parseLine(s string) ([]string, []string) {
	rawData := strings.Fields(s)
	digits := rawData[0:10]
	entry := rawData[11:]
	sort.Slice(digits, func(i, j int) bool {
		return len(digits[i]) < len(digits[j])
	})
	return digits, entry
}

func sortRunesInString(s string) string {
	temp := []rune(s)
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})
	return string(temp)
}

func decipherNumber(numbers []string) map[string]string {
	for i := range numbers {
		numbers[i] = sortRunesInString(numbers[i])
	}

	ret := make(map[string]string)
	decoder := make(map[rune]int)
	for _, v := range numbers {
		for _, c := range v {
			decoder[c]++
		}
	}
	//Determine the Top by picking the only value in 7 that's not in 1
	for _, v := range numbers[1] {
		if !strings.Contains(numbers[0], string(v)) {
			ret["Top"] = string(v)
		}
	}
	//Detemine the bottom by picking the only value that does not appear in 1 4 and 7 and does appear in all 3 numbers that have 5 pieces (2,3,5)
	for _, v := range numbers {
		for _, vv := range v {
			if !strings.Contains(numbers[0], string(vv)) &&
				!strings.Contains(numbers[1], string(vv)) &&
				!strings.Contains(numbers[2], string(vv)) &&
				strings.Contains(numbers[3], string(vv)) &&
				strings.Contains(numbers[4], string(vv)) &&
				strings.Contains(numbers[5], string(vv)) {
				ret["Bottom"] = string(vv)
			}
		}
	}

	for k, v := range decoder {
		switch v {
		case 4:
			ret["BotLeft"] = string(k)
		case 6:
			ret["TopLeft"] = string(k)
		case 7:
			if ret["Bottom"] != string(k) {
				ret["Middle"] = string(k)
			}
		case 8:
			if ret["Top"] != string(k) {
				ret["TopRight"] = string(k)
			}
		case 9:
			ret["BotRight"] = string(k)
		}
	}
	return ret
}

func generateNumberMap(decoder map[string]string) map[string]int {
	staticMap := make(map[string]int)
	zero := sortRunesInString(decoder["Top"] + decoder["TopLeft"] + decoder["TopRight"] + decoder["BotLeft"] + decoder["BotRight"] + decoder["Bottom"])
	staticMap[zero] = 0
	one := sortRunesInString(decoder["TopRight"] + decoder["BotRight"])
	staticMap[one] = 1
	two := sortRunesInString(decoder["Top"] + decoder["TopRight"] + decoder["Middle"] + decoder["BotLeft"] + decoder["Bottom"])
	staticMap[two] = 2
	three := sortRunesInString(decoder["Top"] + decoder["TopRight"] + decoder["Middle"] + decoder["BotRight"] + decoder["Bottom"])
	staticMap[three] = 3
	four := sortRunesInString(decoder["TopLeft"] + decoder["TopRight"] + decoder["Middle"] + decoder["BotRight"])
	staticMap[four] = 4
	five := sortRunesInString(decoder["Top"] + decoder["TopLeft"] + decoder["Middle"] + decoder["BotRight"] + decoder["Bottom"])
	staticMap[five] = 5
	six := sortRunesInString(decoder["Top"] + decoder["TopLeft"] + decoder["Middle"] + decoder["BotLeft"] + decoder["BotRight"] + decoder["Bottom"])
	staticMap[six] = 6
	seven := sortRunesInString(decoder["Top"] + decoder["TopRight"] + decoder["BotRight"])
	staticMap[seven] = 7
	eight := sortRunesInString(decoder["Top"] + decoder["TopLeft"] + decoder["TopRight"] + decoder["Middle"] + decoder["BotLeft"] + decoder["BotRight"] + decoder["Bottom"])
	staticMap[eight] = 8
	nine := sortRunesInString(decoder["Top"] + decoder["TopLeft"] + decoder["TopRight"] + decoder["Middle"] + decoder["BotRight"] + decoder["Bottom"])
	staticMap[nine] = 9
	return staticMap
}

func day8p1(data []string) {
	finalAnswer := 0

	for _, v := range data {
		if len(v) < 1 {
			continue
		}
		_, e := parseLine(v)
		for _, c := range e {
			switch len(c) {
			case 2:
				finalAnswer += 1
			case 3:
				finalAnswer += 1
			case 4:
				finalAnswer += 1
			case 7:
				finalAnswer += 1
			default:
				finalAnswer += 0
			}
		}
	}

	fmt.Println(finalAnswer)
}

func day8p2(data []string) {
	runningTotal := 0
	for _, v := range data {
		if len(v) < 1 {
			continue
		}
		d, e := parseLine(v)
		decoder := decipherNumber(d)
		numbers := generateNumberMap(decoder)
		decodedNumber := 0
		for i := range e {
			e[i] = sortRunesInString(e[i])
			decodedNumber += numbers[e[i]] * int(math.Pow(10, float64(3-i)))
		}
		runningTotal += decodedNumber
	}
	fmt.Println(runningTotal)
}

func main() {
	data, err := aoc.ReadFile(".\\input.txt")
	if err != nil {
		panic(err.Error)
	}

	day8p1(data)
	day8p2(data)
}

//top segment: exists in 7, does not exist in 1 - 8 total
//middle segment: exists in 2/3 of 0 6 9 - 7 total
//bottom segment: exists in every number that isn't 1, 4, 7 - 7 total
//top left segment: exists in 0, 4, 5, 6, 8, 9 - 6 total
//top right segment: exists in 0, 1, 2, 3, 4, 7, 8 ,9 - 8 total
//bottom left segment: exists in 0,2,6,8 - 4 total
//bottom right segment: exists in 0, 1, 3, 4, 5, 6, 7, 8, 9 - 9 total

//2 segments: 1
//3 segments: 7
//4 segments: 4
//5 segments: 2, 3, 5
//6 segments: 0, 6, 9
//7 segments: 8
