package bags

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type bag struct {
	Parents  []string
	Children map[string]int
}

func readFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileBytes), "\n"), nil
}

func populateBag(b string) (string, bag) {
	var thisBag bag
	//name everything consistently. 1 bag 3 bag whatever.
	b = strings.ReplaceAll(b, "bags", "bag")
	//split the string to get our container name.
	split1 := strings.Split(b, " contain ")
	bagname := strings.TrimSpace(split1[0])

	if split1[1] == "no other bag." {
		return bagname, bag{}
	}

	split2 := strings.Split(split1[1], ", ")

	//Strip the trailing period
	split2[len(split2)-1] = split2[len(split2)-1][:len(split2[len(split2)-1])-1]

	thisBag.Children = make(map[string]int, 0)
	for _, i := range split2 {
		num, err := strconv.Atoi(i[0:1])
		if err != nil {
			fmt.Println(i)
			fmt.Println("ERROR")
			continue
		}
		thisBag.Children[strings.TrimSpace(i[1:])] = num
	}
	return bagname, thisBag
}

func main() {
	data, err := readFile("input.txt")
	if err != nil {
		panic("Unable to read data file")
	}

	var bags = make(map[string]bag, 0)

	for _, v := range data {
		if len(v) == 0 {
			continue
		}
		bagname, b := populateBag(v)
		bags[bagname] = b
	}

	var stack []string
	stack = append(stack, "shiny gold bag")
	bagsMet := make(map[string]bool)

	for len(stack) > 0 {
		for k, v := range bags {
			_, ok := v.Children[stack[0]]
			if ok {
				bagsMet[k] = true
				stack = append(stack, k)
			}
		}
		stack[0] = ""
		stack = stack[1:]
	}
	fmt.Println(len(bagsMet))

	var stack2 []string
	stack2 = append(stack2, "shiny gold bag")
	bagsCount := 0
	for len(stack2) > 0 {
		for k, v := range bags[stack2[0]].Children {
			for x := 0; x < v; x++ {
				stack2 = append(stack2, k)
			}
			bagsCount += v
		}
		stack2[0] = ""
		stack2 = stack2[1:]
	}
	fmt.Println(bagsCount)
}
