package main

import (
	"fmt"
)

func testExist(m map[int]int, test int) bool {
	_, ok := m[test]
	if ok {
		return true
	}
	return false
}

func main() {
	n := make(map[int]int)
	n[8] = 1
	n[13] = 2
	n[1] = 3
	n[0] = 4
	n[18] = 5
	n[9] = 6

	next := 0
	//for i := 7; i <= 2020; i++ {
	for i := 7; i <= 30000000; i++ {

		if !testExist(n, next) {
			n[next] = i
			next = 0
		} else {
			temp := i - n[next]
			//fmt.Println(temp)
			n[next] = i
			next = temp
		}
		//if i == 2019 {
		if i == 29999999 {
			fmt.Println("Post:", next)
			//fmt.Println(n)
		}
	}
}
