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

	for _, line := range data {
		num, err := strconv.ParseInt(line, 2, 0)
		if err != nil {
			continue
		}
		if num&0b0000000001 == 1 {
			ParsedBinary[1]++
		}
		if num&0b0000000010 == 2 {
			ParsedBinary[2]++
		}
		if num&0b0000000100 == 4 {
			ParsedBinary[4]++
		}
		if num&0b0000001000 == 8 {
			ParsedBinary[8]++
		}
		if num&0b0000010000 == 16 {
			ParsedBinary[16]++
		}
		if num&0b0000100000 == 32 {
			ParsedBinary[32]++
		}
		if num&0b0001000000 == 64 {
			ParsedBinary[64]++
		}
		if num&0b0010000000 == 128 {
			ParsedBinary[128]++
		}
		if num&0b0100000000 == 256 {
			ParsedBinary[256]++
		}
		if num&0b1000000000 == 512 {
			ParsedBinary[512]++
		}
	}

	gammaRate := 0
	epsilonRate := 0
	for k, v := range ParsedBinary {
		if v > 500 {
			gammaRate += k
		} else {
			epsilonRate += k
		}
	}

	fmt.Println(gammaRate, epsilonRate, gammaRate*epsilonRate)
}
