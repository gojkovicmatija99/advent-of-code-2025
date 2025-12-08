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
	denoms := getDenominators(str)

	if len(str) == 2 && str[0] == str[1] {
		return true
	}

	for _, d := range denoms {
		pnts := make([]int, len(str)/d)
		for i := range len(pnts) {
			pnts[i] = i * d
		}

		for {
			if pnts[0] == d {
				return true
			}

			for i := range len(pnts) - 1 {
				if str[pnts[0]] != str[pnts[i+1]] {
					goto brake
				}
			}

			for i := range len(pnts) {
				pnts[i] = pnts[i] + 1
			}
		}
	brake:
	}

	return false
}

func getDenominators(str string) []int {
	var denoms []int
	for i := 1; i < len(str)-1; i++ {
		if len(str)%i == 0 {
			denoms = append(denoms, i)
		}
	}

	return denoms
}
