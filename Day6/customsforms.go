package customsforms

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	step1 := strings.ReplaceAll(string(fileBytes), "\n\n", ",") //replace form separator with comma
	formattedData := strings.ReplaceAll(step1, "\n", " ")       //replace single newlines with nothing
	return strings.Split(formattedData, ","), nil
}

func countGroup(group string) int {
	check := make(map[rune]bool)
	for _, v := range group {
		if v == ' ' {
			continue
		}
		check[v] = true
	}
	return len(check)
}

func countGroupAll(group string) int {
	check := make(map[rune]int)
	members := strings.Split(group, " ")
	if members[len(members)-1] == "" {
		members = members[:len(members)-1]
	}
	for _, v := range members {
		for _, j := range v {
			check[j]++
		}
	}
	result := 0
	for _, v := range check {
		if v == len(members) {
			result++
		}
	}
	return result
}

func main() {

	forms, err := readFile("input.txt")
	if err != nil {
		panic("Unable to load file")
	}
	runningTotal := 0

	for _, v := range forms {
		runningTotal += countGroup(v)
	}
	fmt.Println("The number of Forms Anyone in a group answered yes to: ", runningTotal)

	runningTotal = 0
	for _, v := range forms {
		runningTotal += countGroupAll(v)
	}
	fmt.Println("The number of Forms Everyone in a group answered yes to: ", runningTotal)

}
