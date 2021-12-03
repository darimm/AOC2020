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
		if num&0b0000000001 == 0b0000000001 {
			ParsedBinary[1]++
		}
		if num&0b0000000010 == 0b0000000010 {
			ParsedBinary[2]++
		}
		if num&0b0000000100 == 0b0000000100 {
			ParsedBinary[4]++
		}
		if num&0b0000001000 == 0b0000001000 {
			ParsedBinary[8]++
		}
		if num&0b0000010000 == 0b0000010000 {
			ParsedBinary[16]++
		}
		if num&0b0000100000 == 0b0000100000 {
			ParsedBinary[32]++
		}
		if num&0b0001000000 == 0b0001000000 {
			ParsedBinary[64]++
		}
		if num&0b0010000000 == 0b0010000000 {
			ParsedBinary[128]++
		}
		if num&0b0100000000 == 0b0100000000 {
			ParsedBinary[256]++
		}
		if num&0b1000000000 == 0b1000000000 {
			ParsedBinary[512]++
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
