package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges := getInput()
	sum := 0
	for _, r := range ranges {
		limits := strings.Split(r, "-")
		sum += invalidNumsSum(limits[0], limits[1])
	}

	fmt.Println(sum)
}

func getInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	return strings.Split(line, ",")
}

func invalidNumsSum(lowStr, highStr string) int {
	low, err := strconv.Atoi(lowStr)
	if err != nil {
		panic(err)
	}

	high, err := strconv.Atoi(highStr)
	if err != nil {
		panic(err)
	}

	sum := 0
	for i := low; i <= high; i++ {
		if isInvalidNum(i) {
			sum += i
		}
	}

	return sum
}

func isInvalidNum(num int) bool {
	str := strconv.Itoa(num)

	if len(str)%2 == 1 {
		return false
	}

	firstIdx := 0
	secondIdx := len(str) / 2
	for {
		if secondIdx == len(str) {
			return true
		}

		if str[firstIdx] != str[secondIdx] {
			return false
		}

		firstIdx++
		secondIdx++
	}
}
