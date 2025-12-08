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
	idx := make([]int, 12)

	for i := 0; i < 12; i++ {
		var startIdx = 0
		if i != 0 {
			startIdx = idx[i-1] + 1
		}

		max := 0
		for j := startIdx; j < len(num)-11+i; j++ {
			if int(num[j]-'0') > max {
				idx[i] = j
				max = int(num[j] - '0')
			}
		}
	}

	res := 0
	for i := 0; i < 12; i++ {
		res += int(num[idx[i]] - '0')
		res *= 10
	}

	return res / 10
}
