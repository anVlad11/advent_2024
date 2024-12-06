package main

import (
	"fmt"
	"github.com/anVlad11/advent_2024/pkg/utils"
	"strings"
	"time"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func do() error {
	input, err := utils.GetInput("inputs/day_06/input_1.txt")
	if err != nil {
		return err
	}

	err = part1(input)
	if err != nil {
		return err
	}

	err = part2(input)
	if err != nil {
		return err
	}

	return nil
}

var icons = map[string][2]int{
	"v": {1, 0},
	"^": {-1, 0},
	">": {0, 1},
	"<": {0, -1},
}

var nextCursor = map[string]string{
	"v": "<",
	"<": "^",
	"^": ">",
	">": "v",
}

const stopIcon = string("#")
const emptyIcon = string(".")
const visitedIcon = string("X")

func printMatrix(matrix [][]string) {
	fmt.Print("\033[H\033[2J")
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%s", matrix[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	time.Sleep(100 * time.Millisecond)
}

func part1(lines []string) error {
	matrix := make([][]string, len(lines)-1) // last line in the input is empty smh

	cursor := [2]int{}

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		matrix[i] = strings.Split(line, "")
		for j := range matrix[i] {
			if _, exists := icons[matrix[i][j]]; exists {
				cursor = [2]int{i, j}
			}
		}
	}

	visitedUnique := 1

	for {
		//printMatrix(matrix)

		modifier := icons[matrix[cursor[0]][cursor[1]]]

		newCursor := [2]int{cursor[0] + modifier[0], cursor[1] + modifier[1]}
		if newCursor[0] < 0 || newCursor[0] > len(matrix)-1 || newCursor[1] < 0 || newCursor[1] > len(matrix[newCursor[0]])-1 {
			matrix[cursor[0]][cursor[1]] = visitedIcon
			break
		}

		if matrix[cursor[0]+modifier[0]][cursor[1]+modifier[1]] == stopIcon {
			matrix[cursor[0]][cursor[1]] = nextCursor[matrix[cursor[0]][cursor[1]]]
			continue
		}

		if matrix[newCursor[0]][newCursor[1]] != visitedIcon {
			visitedUnique++
		}

		matrix[newCursor[0]][newCursor[1]] = matrix[cursor[0]][cursor[1]]
		matrix[cursor[0]][cursor[1]] = visitedIcon
		cursor = newCursor
	}

	fmt.Println(visitedUnique)

	return nil
}

func part2(lines []string) error {
	rawMatrix := make([][]string, len(lines)-1) // last line in the input is empty smh

	startingCursorPosition := [2]int{}
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		rawMatrix[i] = strings.Split(line, "")
		for j := range rawMatrix[i] {
			if _, exists := icons[rawMatrix[i][j]]; exists {
				startingCursorPosition = [2]int{i, j}
			}
		}
	}

	matrix := copyMatrix(rawMatrix)

	visitedWithDirections, _ := walkMatrix(startingCursorPosition, matrix)

	loopsAmount := 0

	for visitedPosition, m := range visitedWithDirections {
		if visitedPosition == startingCursorPosition {
			continue
		}
		isLoop := false
		for _, stepsVisited := range m {
			for range stepsVisited {
				matrix = copyMatrix(rawMatrix)
				matrix[visitedPosition[0]][visitedPosition[1]] = stopIcon

				_, isLoop = walkMatrix(startingCursorPosition, matrix)
				if isLoop {
					loopsAmount++
				}

				if isLoop {
					break
				}
			}
			if isLoop {
				break
			}
		}
	}

	fmt.Println(loopsAmount)

	return nil
}

func walkMatrix(startingCursorPosition [2]int, matrix [][]string) (map[[2]int]map[string][]int64, bool) {
	cursorPosition := [2]int{startingCursorPosition[0], startingCursorPosition[1]}
	// coordinates of each position visited with the direction,
	// including steps at which it was visited with this direction
	visitedWithDirections := map[[2]int]map[string][]int64{}
	step := int64(0)

	for ; ; step++ {

		cursorIcon := matrix[cursorPosition[0]][cursorPosition[1]]

		if _, exists := visitedWithDirections[cursorPosition]; !exists {
			visitedWithDirections[cursorPosition] = map[string][]int64{}
		}
		visitedWithDirections[cursorPosition][cursorIcon] = append(visitedWithDirections[cursorPosition][cursorIcon], step)
		if len(visitedWithDirections[cursorPosition][cursorIcon]) > 1 {
			// path is a loop, let's return early
			return visitedWithDirections, true
		}

		cursorPositionModifier := icons[cursorIcon]

		newCursorPosition := [2]int{cursorPosition[0] + cursorPositionModifier[0], cursorPosition[1] + cursorPositionModifier[1]}
		if newCursorPosition[0] < 0 || newCursorPosition[0] > len(matrix)-1 || newCursorPosition[1] < 0 || newCursorPosition[1] > len(matrix[newCursorPosition[0]])-1 {
			matrix[cursorPosition[0]][cursorPosition[1]] = visitedIcon
			break
		}

		if matrix[newCursorPosition[0]][newCursorPosition[1]] == stopIcon {
			matrix[cursorPosition[0]][cursorPosition[1]] = nextCursor[cursorIcon]
			continue
		}

		matrix[newCursorPosition[0]][newCursorPosition[1]] = matrix[cursorPosition[0]][cursorPosition[1]]
		matrix[cursorPosition[0]][cursorPosition[1]] = visitedIcon
		cursorPosition = newCursorPosition
	}

	return visitedWithDirections, false
}

func copyMatrix(matrix [][]string) [][]string {
	newMatrix := make([][]string, len(matrix))
	for i := range matrix {
		newMatrix[i] = make([]string, len(matrix[i]))
		copy(newMatrix[i], matrix[i])
	}

	return newMatrix
}
