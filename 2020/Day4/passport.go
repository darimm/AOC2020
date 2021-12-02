package passport

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type passport map[string]string

func readFile(fileName string) ([]passport, error) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	step1 := strings.Replace(string(fileBytes), "\n\n", ",", -1) //replace passport separator with comma
	formattedData := strings.Replace(step1, "\n", " ", -1)       //replace single newlines with a space for consistency
	listPassports := strings.Split(formattedData, ",")
	passports := make([]passport, 0)

	for _, i := range listPassports {
		entry := make(passport, 0)
		passParts := strings.Split(i, " ")
		for _, j := range passParts {
			kv := strings.Split(j, ":")
			if len(kv) == 2 {
				entry[kv[0]] = kv[1]
			}
		}
		passports = append(passports, entry)
	}

	return passports, nil
}

func isBetween(x, low, high int) bool {
	if x < low || x > high {
		return false
	}
	return true
}

func validateByr(byr string) bool {
	year, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}
	if !isBetween(year, 1920, 2020) {
		return false
	}
	return true
}

func validateIyr(iyr string) bool {
	year, err := strconv.Atoi(iyr)
	if err != nil {
		return false
	}
	if !isBetween(year, 2010, 2020) {
		return false
	}
	return true
}

func validateEyr(eyr string) bool {
	year, err := strconv.Atoi(eyr)
	if err != nil {
		return false
	}
	if !isBetween(year, 2020, 2030) {
		return false
	}
	return true
}

func validateHgt(hgt string) bool {
	if len(hgt) < 4 {
		return false
	}
	switch units := hgt[len(hgt)-2:]; units {
	case "cm":
		test, err := strconv.Atoi(strings.Split(hgt, "cm")[0])
		if err != nil {
			return false
		}
		if test < 150 || test > 193 {
			return false
		}
	case "in":
		test, err := strconv.Atoi(strings.Split(hgt, "in")[0])
		if err != nil {
			return false
		}
		if test < 59 || test > 76 {
			return false
		}
	default:
		return false
	}
	return true
}

func validateHcl(hcl string) bool {
	if len(hcl) != 7 {
		return false
	}
	_, err := hex.DecodeString(hcl[1:])
	if err != nil {
		return false
	}
	return true
}

func validateEcl(ecl string) bool {
	testValues := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	_, ok := testValues[ecl]
	if !ok {
		return false
	}
	return true
}

func validatePid(pid string) bool {
	if len(pid) != 9 {
		return false
	}
	_, err := strconv.Atoi(pid)
	if err != nil {
		return false
	}
	return true
}

func validatePassport(v map[string]string) bool {
	if validateByr(v["byr"]) && validateIyr(v["iyr"]) && validateEyr(v["eyr"]) && validateHgt(v["hgt"]) && validateHcl(v["hcl"]) && validateEcl(v["ecl"]) &&
		validatePid(v["pid"]) {
		return true
	}
	return false
}

func main() {
	passportList, err := readFile("input.txt")
	if err != nil {
		panic("Unable to read file")
	}

	validPassports := 0
	for _, i := range passportList {
		if validatePassport(i) {
			validPassports++
		}
	}
	fmt.Println(validPassports)

}
