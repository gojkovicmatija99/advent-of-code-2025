package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var START_POSITION = 50

func main() {
	zero := 0
	currPos := START_POSITION

	lines := getInput()
	for _, l := range lines {
		num := rotationToNumber(l)
		currPos = currPos + num
		zero += numOfTouchedZero(currPos-num, currPos)
		currPos = currPos % 100
	}
	fmt.Println(zero)
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

func rotationToNumber(rot string) int {
	dir := rot[0]
	num, err := strconv.Atoi(rot[1:])
	if err != nil {
		panic(err)
	}

	if dir == 'L' {
		return num * -1
	} else {
		return num
	}
}

func numOfTouchedZero(oldPos, currPos int) int {
	zero := int(math.Abs(float64(currPos) / 100.0))
	if currPos == 0 || (currPos > 0 && oldPos < 0) || (currPos < 0 && oldPos > 0) {
		zero++
	}

	return zero
}
