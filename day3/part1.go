package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := getInput()
	sum := 0
	for _, l := range lines {
		sum += getLargestJoltage(l)
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

func getLargestJoltage(num string) int {
	oneIdx := len(num) - 1
	tenIdx := len(num) - 2
	for i := len(num) - 3; i >= 0; i-- {
		if num[i] >= num[tenIdx] {
			if num[tenIdx] > num[oneIdx] {
				oneIdx = tenIdx
			}
			tenIdx = i
		}
	}

	return int((num[tenIdx]-'0')*10) + int((num[oneIdx] - '0'))
}
