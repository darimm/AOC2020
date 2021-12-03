package main

import (
	"fmt"
	"io/ioutil"
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

func main() {
	data, err := readFile(".\\input.txt")
	if err != nil {
		panic(err.Error)
	}

	ParsedBinary := make(map[int]int)
	ParsedBinary[1] = 0
	ParsedBinary[2] = 0
	ParsedBinary[4] = 0
	ParsedBinary[8] = 0
	ParsedBinary[16] = 0
	ParsedBinary[32] = 0
	ParsedBinary[64] = 0
	ParsedBinary[128] = 0
	ParsedBinary[256] = 0
	ParsedBinary[512] = 0
	ParsedBinary[1024] = 0
	ParsedBinary[2048] = 0

	for _, line := range data {
		num, err := strconv.ParseInt(line, 2, 0)
		if err != nil {
			continue
		}
		if num&1 == 1 {
			ParsedBinary[1]++
		}
		if num&2 == 2 {
			ParsedBinary[2]++
		}
		if num&4 == 4 {
			ParsedBinary[4]++
		}
		if num&8 == 8 {
			ParsedBinary[8]++
		}
		if num&16 == 16 {
			ParsedBinary[16]++
		}
		if num&32 == 32 {
			ParsedBinary[32]++
		}
		if num&64 == 64 {
			ParsedBinary[64]++
		}
		if num&128 == 128 {
			ParsedBinary[128]++
		}
		if num&256 == 256 {
			ParsedBinary[256]++
		}
		if num&512 == 512 {
			ParsedBinary[512]++
		}
		if num&1024 == 1024 {
			ParsedBinary[1024]++
		}
		if num&2048 == 2048 {
			ParsedBinary[2048]++
		}
	}

	gammaRate := 0
	epsilonRate := 0
	for k, v := range ParsedBinary {
		if v > (len(data)-1)/2 {
			fmt.Printf("Adding %d to Gamma\r\n", k)
			gammaRate += k
		} else {
			fmt.Printf("Adding %d to Epsilon\r\n", k)
			epsilonRate += k
		}
	}

	fmt.Println(gammaRate, epsilonRate, gammaRate*epsilonRate)
	fmt.Println(ParsedBinary)
}
