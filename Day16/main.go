package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	min1, max1, min2, max2 int
}

func readFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileBytes), "\n\n"), nil
}

func makeRules(lines string) map[string]rule {
	returnVal := make(map[string]rule)
	r := regexp.MustCompile(`^(.*): (\d+)-(\d+) or (\d+)-(\d+)$`)
	rules := strings.Split(lines, "\n")

	for _, s := range rules {
		parts := r.FindStringSubmatch(s)
		if len(parts) != 6 {
			fmt.Println("Error Occurred: ", parts)
			continue
		}
		var rVals []int
		for i := 2; i <= 5; i++ {
			m, err := strconv.Atoi(parts[i])
			if err != nil {
				fmt.Println("Error converting Value to int: ", err, parts[i])
			}
			rVals = append(rVals, m)
		}
		rv := rule{
			min1: rVals[0],
			max1: rVals[1],
			min2: rVals[2],
			max2: rVals[3],
		}
		returnVal[parts[1]] = rv
	}
	return returnVal
}

func parseTicket(t string) map[string]int {
	s := strings.Split(t, ",")
	if len(s) != 20 {
		fmt.Println("String does not contain appropriate number of values: ", s)
	}
	returnVal := make(map[string]int)
	for k, v := range s {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Non integer value found")
		}
		returnVal[fmt.Sprintf("%d", k)] = num
	}
	return returnVal
}

func testRulesP1(ticket map[string]int, rules map[string]rule) int {
	var returnVal int
	for _, v := range ticket {
		test := false
		for _, i := range rules {
			if (v >= i.min1 && v <= i.max1) || (v >= i.min2 && v <= i.max2) {
				test = true
			}
		}
		if !test {
			returnVal += v
		}
	}
	return returnVal
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func testRulesP2(tickets []map[string]int, rules map[string]rule) map[int]string {
	returnVal := make(map[int]string)
	possiblePositions := make(map[string][]int)
	var tickets2 []map[string]int
	for _, i := range tickets {
		if testRulesP1(i, rules) == 0 {
			tickets2 = append(tickets2, i)
		}
	}

	for ruleName, thisRule := range rules {
		for testTickets := 0; testTickets <= 19; /*Magic Number. It's how many fields are in each ticket*/ testTickets++ {
			//fmt.Printf("Testing Ticket Position %d against rule %s %+v\n", testTickets, ruleName, thisRule)
			myTest := true
			for _, ticket := range tickets2 {
				index := fmt.Sprintf("%d", testTickets)
				//fmt.Println(index, ticket[index])
				if !((ticket[index] >= thisRule.min1 && ticket[index] <= thisRule.max1) || (ticket[index] >= thisRule.min2 && ticket[index] <= thisRule.max2)) {
					myTest = false
					break
				}
			}
			if myTest {
				possiblePositions[ruleName] = append(possiblePositions[ruleName], testTickets)
			}
		}
	}
	for {
		for k, v := range possiblePositions {
			if len(v) == 1 {
				returnVal[v[0]] = k
				for j := range possiblePositions {
					if len(possiblePositions[j]) > 1 {
						for i := 0; i < len(possiblePositions[j]); i++ {
							if possiblePositions[j][i] == v[0] {
								possiblePositions[j] = remove(possiblePositions[j], i)
							}
						}

					}
				}
			}
		}
		if len(returnVal) == 20 {
			break
		}
	}
	return returnVal
}

func main() {
	lines, err := readFile("input.txt")
	if err != nil {
		panic("Unable to load File")
	}
	rules := makeRules(lines[0])

	myTicket := parseTicket(strings.Split(lines[1], "\n")[1])

	var tickets []map[string]int

	for _, v := range strings.Split(lines[2], "\n") {
		if v == "nearby tickets:" || v == "" {
			continue
		}
		tickets = append(tickets, parseTicket(v))
	}
	tickets = append(tickets, myTicket)
	accumulator := 0
	for _, testTicket := range tickets {
		accumulator += testRulesP1(testTicket, rules)
	}
	fmt.Println(accumulator)

	//testRulesP2(tickets, rules)
	ruleNames := testRulesP2(tickets, rules)
	fmt.Println(ruleNames)
	p2return := 1
	for k, v := range ruleNames {
		if strings.HasPrefix(v, "departure") {
			p2return *= myTicket[fmt.Sprintf("%d", k)]
		}
	}
	fmt.Println(p2return)
}
