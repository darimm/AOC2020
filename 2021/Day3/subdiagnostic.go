package subdiagnostic

import (
	"fmt"
	"io/ioutil"
	"math"
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

func countBits(data []string) map[int]int {
	ParsedBinary := make(map[int]int)
	for mapMake := 1; mapMake <= 2048; mapMake = mapMake << 1 {
		ParsedBinary[mapMake] = 0
	}

	for _, line := range data {
		num, err := strconv.ParseInt(line, 2, 0)
		if err != nil {
			continue
		}

		for testVal := int64(1); testVal <= 2048; testVal = testVal << 1 {
			if num&testVal == testVal {
				ParsedBinary[int(testVal)]++
			}
		}
	}
	return ParsedBinary
}

func CalculateGammaEpsilon(bitsMap map[int]int, dataLength int) (gammaRate, epsilonRate int) {

	gammaRate = 0
	epsilonRate = 0
	for k, v := range bitsMap {
		if v > (dataLength-1)/2 {
			gammaRate += k
		} else {
			epsilonRate += k
		}
	}
	return gammaRate, epsilonRate
}

func filterSlice(data []string, testValue int, comparitor func(int, int) bool) []string {
	dataCopy := make([]string, 0)
	for k, i := range data {
		num, err := strconv.ParseInt(i, 2, 0)
		if err != nil {
			continue
		}
		if comparitor(int(num), testValue) {
			dataCopy = append(dataCopy, data[k])
		}
	}
	return dataCopy
}

func compareAnd(val, mask int) bool {
	if val&mask == mask {
		return true
	}
	return false
}

func compareXor(val, mask int) bool {
	if val^mask > val {
		return true
	}
	return false
}

func OxygenFilter(data []string) int64 {
	dataCopy := make([]string, 0)
	dataCopy = append(dataCopy, data...)

	for len(dataCopy) > 1 {
		for i := 2048; i > 0; i = i >> 1 {
			temp := countBits(dataCopy)
			divisor := math.Ceil((float64(len(dataCopy)) / 2))
			if len(dataCopy) == 1 {
				break
			}
			if temp[i] >= int(divisor) {
				dataCopy = filterSlice(dataCopy, i, compareAnd)
			} else {
				dataCopy = filterSlice(dataCopy, i, compareXor)
			}
		}
	}
	fmt.Println("Final Oxygen Filter Value: ", dataCopy)
	num, err := strconv.ParseInt(dataCopy[0], 2, 0)
	if err != nil {
		panic("This should definitely never happen")
	}
	return num
}

func CO2Filter(data []string) int64 {
	dataCopy := make([]string, 0)
	dataCopy = append(dataCopy, data...)

	for len(dataCopy) > 1 {
		for i := 2048; i > 0; i = i >> 1 {
			temp := countBits(dataCopy)
			divisor := math.Ceil((float64(len(dataCopy)) / 2))
			if len(dataCopy) == 1 {
				break
			}
			if temp[i] < int(divisor) {
				dataCopy = filterSlice(dataCopy, i, compareAnd)
			} else {
				dataCopy = filterSlice(dataCopy, i, compareXor)
			}
		}
	}
	fmt.Println("Final CO2 Filter Value: ", dataCopy)
	num, err := strconv.ParseInt(dataCopy[0], 2, 0)
	if err != nil {
		panic("This should definitely never happen")
	}
	return num
}

func main() {
	data, err := readFile(".\\input.txt")
	if err != nil {
		panic(err.Error)
	}

	bitsMap := countBits(data)
	gammaRate, epsilonRate := CalculateGammaEpsilon(bitsMap, len(data))

	fmt.Println(gammaRate, epsilonRate, gammaRate*epsilonRate)
	fmt.Println(bitsMap)

	Oxy := OxygenFilter(data[:len(data)-1])
	CO2 := CO2Filter(data[:len(data)-1])
	fmt.Println(Oxy, CO2)
	fmt.Println(Oxy * CO2)
}
