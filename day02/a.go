package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_input() []string {
	var lines []string
	readFile, err := os.Open("./input.txt")
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}

var choiceToNum = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
	"A": 1,
	"B": 2,
	"C": 3,
}

func gameScore(opponentChoice int, myChoice int) int {
	if myChoice == 1 && opponentChoice == 3 {
		return 6
	} else if myChoice-opponentChoice == 1 {
		return 6
	} else if myChoice == opponentChoice {
		return 3
	} else {
		return 0
	}
}

func main() {
	input := get_input()
	score := 0
	for _, line := range input {
		choices := strings.Split(line, " ")
		opponentChoice := choiceToNum[choices[0]]
		myChoice := choiceToNum[choices[1]]
		score += myChoice + gameScore(opponentChoice, myChoice)
	}
	fmt.Println(score)
}
