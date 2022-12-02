package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func main() {
	input := get_input()
	maxCalories := 0
	elfCalories := 0
	for _, line := range input {
		if line == "" {
			elfCalories = 0
		} else {
			calorie, err := strconv.Atoi(line)
			check(err)
			elfCalories += calorie
		}
		if elfCalories > maxCalories {
			maxCalories = elfCalories
		}
	}
	fmt.Println(maxCalories)
}
