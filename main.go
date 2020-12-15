package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type mask [36]int
type register [36]bool

func readFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileBytes), "\n"), nil
}

func (m *mask) parseLine(line string, registerList map[int]register) bool {
	s := strings.Split(line, " = ")
	//If we're updating the subnet mask make m[0] the 2^0 bit
	if s[0] == "mask" {
		for i := 35; i >= 0; i-- {
			switch s[1][i] {
			case '0':
				m[35-i] = 0
			case '1':
				m[35-i] = 1
			case 'X':
				m[35-i] = -1
			}
		}
		return false
	}
	r := regexp.MustCompile(`^mem\[(\d+)\]$`)
	rval, err := strconv.Atoi(r.FindStringSubmatch(s[0])[1])
	if err != nil {
		fmt.Println("Unable to convert register number to int")
	}

	bignum, err := strconv.ParseInt(s[1], 10, 36)
	if err != nil {
		fmt.Println("Error occurred converting bignum", err)
	}

	bignumBase2 := strconv.FormatInt(bignum, 2)
	returnVal := [36]bool{}

	for i := len(bignumBase2) - 1; i >= 0; i-- {
		pos := len(bignumBase2) - 1 - i
		switch m[pos] {
		case 1:
			returnVal[pos] = true
		case 0:
			returnVal[pos] = false
		case -1:
			rv, err := strconv.Atoi(string(bignumBase2[i]))
			if err != nil {
				fmt.Println("ERROR: ", err)
			}
			if rv == 0 {
				returnVal[pos] = false || registerList[rval][pos]
			} else {
				returnVal[pos] = true
			}
		}
	}
	registerList[rval] = returnVal
	return true
}

// Pow returns x**y, the base-x exponential of y.
func Pow(x, y int64) (r int64) {
	if x == r {
		return
	}
	r = 1
	if x == r {
		return
	}
	for y > 0 {
		if y&1 == 1 {
			r *= x
		}
		x *= x
		y >>= 1
	}
	return
}

func registerToBigInt(r register) int64 {
	returnVal := int64(0)
	for pos := 0; pos < len(r); pos++ {
		if r[pos] {
			returnVal += Pow(2, int64(pos))
		}
	}
	return returnVal
}

func main() {
	lines, _ := readFile("input.txt")
	registers := make(map[int]register)
	var bitmask mask
	for _, line := range lines {
		if line == "" {
			continue
		}
		bitmask.parseLine(line, registers)
	}
	var finalVal int64
	for _, i := range registers {
		//fmt.Println(k, i)
		finalVal += registerToBigInt(i)
	}
	fmt.Println(finalVal)
}
