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
func charToNum(char rune) int {
	if char > 64 && char < 91 {
		return int(char - 64 + 26)
	} else {
		return int(char - 96)
	}
}

func main() {
	input := get_input()
	answer := 0
	group := []string{}
	for _, line := range input {
		if len(group) < 3 {
			group = append(group, line)
		} else {
			group = []string{line}
		}
		if len(group) == 3 {
			for _, char := range group[0] {
				if strings.ContainsRune(group[1], char) && strings.ContainsRune(group[2], char) {
					answer += charToNum(char)
					break
				}
			}
		}
	}
	fmt.Println(answer)
}
