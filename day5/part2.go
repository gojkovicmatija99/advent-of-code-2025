package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ranges, _ := getInput()

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	mergedRangeIdx := 0
	writeIdx := 0
	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] >= ranges[mergedRangeIdx][0] && ranges[i][0] <= ranges[mergedRangeIdx][1] {
			ranges[mergedRangeIdx][1] = int(math.Max(float64(ranges[i][1]), float64(ranges[mergedRangeIdx][1])))
		} else {
			ranges[writeIdx] = ranges[mergedRangeIdx]
			writeIdx++
			mergedRangeIdx = i
		}
	}

	ranges[writeIdx] = ranges[mergedRangeIdx]
	writeIdx++
	ranges = ranges[:writeIdx]

	fresh := 0
	for i := 0; i < len(ranges); i++ {
		fresh += ranges[i][1] - ranges[i][0] + 1
	}

	fmt.Println(fresh)
}

func getInput() ([][]int, []int) {

	var ranges [][]int
	var points []int

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	readingRanges := true

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			readingRanges = false
			continue
		}

		if readingRanges {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				continue
			}
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, []int{a, b})
		} else {
			val, _ := strconv.Atoi(line)
			points = append(points, val)
		}
	}

	return ranges, points
}
