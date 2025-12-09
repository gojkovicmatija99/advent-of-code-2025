package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := getInput()

	sum := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			sum += access(lines, i, j)

		}
	}
	fmt.Println(sum)
}

func getInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func access(lines []string, row, col int) int {
	if lines[row][col] != '@' {
		return 0
	}

	rowMove := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	colMove := []int{-1, 0, 1, 1, 1, 0, -1, -1}

	count := 0
	for i := 0; i < 8; i++ {
		lookRow := rowMove[i] + row
		lookCol := colMove[i] + col

		if lookRow < 0 || lookCol < 0 || lookRow >= len(lines) || lookCol >= len(lines[0]) {
			continue
		}

		if lines[lookRow][lookCol] == '@' {
			count++
		}
	}

	if count < 4 {
		return 1
	} else {
		return 0
	}
}
