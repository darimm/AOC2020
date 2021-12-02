package pretendassembly

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instructions struct {
	step        int
	accumulator int
	sequence    []instruction
}

type instruction struct {
	asm string
	val int
}

func (i *instructions) nop() {
	i.step++
}

func (i *instructions) acc() {
	i.accumulator += i.sequence[i.step].val
	i.step++
}

func (i *instructions) jmp() {
	i.step += i.sequence[i.step].val
}

func (i *instructions) executeCommand(step int) {
	switch i.sequence[step].asm {
	case "nop":
		i.nop()
	case "acc":
		i.acc()
	case "jmp":
		i.jmp()
	}
}

func (i *instructions) findBadInstructionAccumulator() bool {
	visited := make(map[int]bool)
	for {
		//fmt.Printf("%v %v %v", i.step, len(i.sequence), i.sequence[i.step])
		if i.step > len(i.sequence)-1 {
			return true
		}
		visited[i.step] = true
		i.executeCommand(i.step)
		if visited[i.step] {
			return false
		}
	}
}

func readFile(fileName string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileBytes), "\n"), nil
}

func parseBadInstructions(data []string) []instructions {
	var returnVal []instructions

	for k, i := range data {
		if strings.HasPrefix(i, "nop") {
			d := append([]string(nil), data...)
			/*d := make([]string, len(data))
			_ = copy(d, data)*/
			d[k] = strings.ReplaceAll(d[k], "nop", "jmp")
			returnVal = append(returnVal, parseInstructions(d))
			d = []string{}
		}
		if strings.HasPrefix(i, "jmp") {
			d := append([]string(nil), data...)
			/*d := make([]string, len(data))
			_ = copy(d, data)*/
			d[k] = strings.ReplaceAll(d[k], "jmp", "nop")
			returnVal = append(returnVal, parseInstructions(d))
			d = []string{}
		}

	}
	return returnVal
}

func parseInstructions(data []string) instructions {
	var returnVal instructions
	var inst []instruction
	for _, i := range data {
		//stupid blank line
		if len(i) == 0 {
			continue
		}
		s := strings.Split(i, " ")
		v, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println("ERROR")
		}
		inst = append(inst, instruction{asm: s[0], val: v})
	}
	returnVal.sequence = inst
	return returnVal
}

func main() {
	data, err := readFile("input.txt")
	if err != nil {
		panic("Unable to load data")
	}

	programData := parseBadInstructions(data)
	for i := range programData {
		if programData[i].findBadInstructionAccumulator() {
			fmt.Println("NO LOOP:", programData[i].accumulator)
			break
		}
	}

	/*programData := parseInstructions(data)
	programData.findBadInstructionAccumulator()
	fmt.Println(programData.accumulator)*/
}
