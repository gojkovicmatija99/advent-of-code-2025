package part1

import (
	"bufio"
	"fmt"
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
		if currPos%100 == 0 {
			zero++
		}
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
	num = num % 100

	if dir == 'L' {
		num = 100 - num
	}

	return num
}
