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

type bingoCard [][]string

func readBingo(raw []string) bingoCard {
	ret := make(bingoCard, 5)
	for i := 0; i < 5; i++ {
		temp := strings.Fields(raw[i])
		ret[i] = append(ret[i], temp...)
	}
	return ret
}

func populateBingoCards(data []string) []bingoCard {
	cardStart := 2
	cardEnd := 0
	var cards []bingoCard
	for i := 2; i < len(data); i++ {
		if data[i] == "" {
			cardEnd = i
			cards = append(cards, readBingo(data[cardStart:cardEnd]))
			cardStart = i + 1
		}
	}
	return cards
}

func playNumber(cards []bingoCard, number string) []bingoCard {
	for _, card := range cards {
	Eachcard:
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				if card[x][y] == number {
					card[x][y] = "X"
					break Eachcard
				}
			}
		}
	}
	return cards
}

func checkWin(cards []bingoCard) (bingoCard, error) {
	for _, card := range cards {
		for y := 0; y < 5; y++ {
			if checkRowOrColumn(card[y]) {
				return card, nil
			}
			column := []string{card[0][y], card[1][y], card[2][y], card[3][y], card[4][y]}
			if checkRowOrColumn(column) {
				return card, nil
			}
		}
	}
	return nil, fmt.Errorf("Not a winner")
}

func checkWinp2(cards []bingoCard) (int, error) {
	for idx, card := range cards {
		for y := 0; y < 5; y++ {
			if checkRowOrColumn(card[y]) {
				return idx, nil
			}
			column := []string{card[0][y], card[1][y], card[2][y], card[3][y], card[4][y]}
			if checkRowOrColumn(column) {
				return idx, nil
			}
		}
	}
	return -1, fmt.Errorf("Not a winner")
}

func checkRowOrColumn(row []string) bool {
	for _, x := range row {
		if x != "X" {
			return false
		}
	}
	return true
}

func day4p1(data []string) {
	bingoNumbers := strings.Split(data[0], ",")
	var winningCard bingoCard
	var winningNumber int
	var winningScore int

	cards := populateBingoCards(data)

	for x := 0; x < len(bingoNumbers); x++ {
		cards = playNumber(cards, bingoNumbers[x])
		test, err := checkWin(cards)
		if err == nil {
			winningCard = test
			winningNumber, _ = strconv.Atoi(bingoNumbers[x])
			break
		}
	}

	for _, row := range winningCard {
		for _, col := range row {
			if col != "X" {
				score, _ := strconv.Atoi(col)
				winningScore += score
			}
		}
	}
	fmt.Println("Winning Score of first card to win: ", winningNumber*winningScore)
}

func day4p2(data []string) {
	bingoNumbers := strings.Split(data[0], ",")
	var winningCard bingoCard
	var winningNumber int
	var winningScore int

	cards := populateBingoCards(data)

	for x := 0; x < len(bingoNumbers); x++ {
		cards = playNumber(cards, bingoNumbers[x])
	retestWin:
		test, err := checkWinp2(cards)
		if len(cards) == 1 && err == nil {
			winningNumber, _ = strconv.Atoi(bingoNumbers[x])
			winningCard = cards[0]
			break
		}
		if err == nil {
			winningNumber, _ = strconv.Atoi(bingoNumbers[x])
			cards = append(cards[:test], cards[test+1:]...)
			goto retestWin
		}
	}

	for _, row := range winningCard {
		for _, col := range row {
			if col != "X" {
				score, _ := strconv.Atoi(col)
				winningScore += score
			}
		}
	}
	fmt.Println("Winning Score of final card to win: ", winningNumber*winningScore)
}

func main() {
	data, err := readFile(".\\input.txt")
	if err != nil {
		panic(err.Error)
	}

	day4p1(data)
	day4p2(data)

}
