package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := getInput()

	graphToZero := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		graphToZero[i] = make([]int, len(lines[0]))
		for j := 0; j < len(lines[i]); j++ {
			graphToZero[i][j] = neighbourToDeleteToZero(lines, i, j)
		}
	}

	checked := make([][]bool, len(lines))
	for i := 0; i < len(lines); i++ {
		checked[i] = make([]bool, len(lines[0]))
	}

	var bfsQueue [][]int
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if graphToZero[i][j] >= 0 && graphToZero[i][j] < 4 {
				bfsQueue = append(bfsQueue, []int{i, j})
			}
		}
	}

	bfs(bfsQueue, checked, graphToZero)

	zero := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if graphToZero[i][j] == 0 {
				zero++
			}
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

func neighbourToDeleteToZero(lines []string, row, col int) int {
	if lines[row][col] == '.' {
		return -1
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

	return count
}

// go bfs from starting cell
// if current cell can't be removed, add adjesent not checked cells to queue
// if current cell can be removed, remove it, decrement adjesent cells and add them to queue

func bfs(bfsQueue [][]int, checked [][]bool, graphToZero [][]int) {
	for len(bfsQueue) > 0 {
		node := bfsQueue[0]
		bfsQueue = bfsQueue[1:]
		row := node[0]
		col := node[1]

		if graphToZero[row][col] <= 0 {
			continue
		}

		if graphToZero[row][col] < 4 {
			graphToZero[row][col] = 0
		}

		rowMove := []int{-1, -1, -1, 0, 1, 1, 1, 0}
		colMove := []int{-1, 0, 1, 1, 1, 0, -1, -1}
		for i := 0; i < 8; i++ {
			lookRow := rowMove[i] + row
			lookCol := colMove[i] + col

			if lookRow < 0 || lookCol < 0 || lookRow >= len(graphToZero) || lookCol >= len(graphToZero[0]) {
				continue
			}

			if graphToZero[lookRow][lookCol] <= 0 {
				continue
			}

			if graphToZero[row][col] == 0 {
				graphToZero[lookRow][lookCol] -= 1
			}

			if graphToZero[row][col] == 0 || !checked[lookRow][lookCol] {
				bfsQueue = append(bfsQueue, []int{lookRow, lookCol})
				checked[lookRow][lookCol] = true
			}
		}
	}
}
