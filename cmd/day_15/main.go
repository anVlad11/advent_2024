package main

import (
	"fmt"
	"github.com/anVlad11/advent_2024/pkg/utils"
	"strings"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func do() error {
	input, err := utils.GetInput("inputs/day_15/input.txt")
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

var directions = map[string][2]int{
	"^": {-1, 0},
	">": {0, 1},
	"v": {1, 0},
	"<": {0, -1},
}

func part1(lines []string) error {
	matrix := [][]string{}

	pos := [2]int{}
	breakLine := 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			breakLine = i
			break
		}
		line := make([]string, len(lines[i]))
		for j := range lines[i] {
			v := string(lines[i][j])
			if v == "@" {
				pos = [2]int{i, j}
			}
			line[j] = v
		}
		matrix = append(matrix, line)
	}

	for i := breakLine; i < len(lines); i++ {
		for j := range lines[i] {
			v := string(lines[i][j])
			posesToMove := [][2]int{pos}
			nextP := pos
			for {
				nextP = [2]int{nextP[0] + directions[v][0], nextP[1] + directions[v][1]}
				if nextP[0] < 0 || nextP[0] >= len(matrix) || nextP[1] < 0 || nextP[1] >= len(matrix[nextP[0]]) {
					posesToMove = [][2]int{}
					break
				}

				done := false
				switch matrix[nextP[0]][nextP[1]] {
				case "O":
					posesToMove = append(posesToMove, nextP)
				case "#":
					posesToMove = [][2]int{}
					done = true
				case ".":
					posesToMove = append(posesToMove, nextP)
					done = true
				}

				if done {
					break
				}
			}

			for k := len(posesToMove) - 1; k > 0; k-- {
				to := [2]int{posesToMove[k][0], posesToMove[k][1]}
				from := [2]int{posesToMove[k-1][0], posesToMove[k-1][1]}
				matrix[to[0]][to[1]] = matrix[from[0]][from[1]]
				matrix[from[0]][from[1]] = "."
				if matrix[to[0]][to[1]] == "@" {
					pos = to
				}
			}

		}
	}

	sum := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != "O" {
				continue
			}
			sum += (100 * i) + j
		}
	}

	fmt.Println(sum)

	return nil
}

func part2(lines []string) error {
	matrix := [][]string{}

	robotPos := [2]int{}
	breakLine := 0
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			breakLine = i
			break
		}

		line := make([]string, len(lines[i])*2)
		for j := range lines[i] {
			v := string(lines[i][j])
			switch v {
			case "O":
				line[j*2] = "["
				line[j*2+1] = "]"
			case "@":
				robotPos = [2]int{i, j * 2}
				line[j*2] = "@"
				line[j*2+1] = "."
			default:
				line[j*2] = v
				line[j*2+1] = v
			}
		}

		matrix = append(matrix, line)
	}

	//fmt.Printf("Input: \n")
	//matrixStr := sprintMatrix(matrix)
	//fmt.Printf("%s\n", matrixStr)
	//fmt.Printf("\n")

	for i := breakLine; i < len(lines); i++ {
		for j := range lines[i] {
			v := string(lines[i][j])
			posesToMove := [][2]int{}
			nextPoses := map[[2]int][2]int{}
			uniquePosesToMove := map[[2]int]struct{}{}
			queue := [][2]int{robotPos}
			for len(queue) > 0 {
				item := queue[0]
				queue = queue[1:]

				if item[0] < 0 || item[0] >= len(matrix) || item[1] < 0 || item[1] >= len(matrix[item[0]]) {
					posesToMove = [][2]int{}
					break
				}

				done := false
				targetV := matrix[item[0]][item[1]]
				switch targetV {
				case "[":
					posesToMove = append(posesToMove, item)

					nextPosQueue := [2]int{item[0] + directions[v][0], item[1] + directions[v][1]}
					nextPoses[item] = nextPosQueue

					if _, ok := uniquePosesToMove[nextPosQueue]; !ok {
						queue = append(queue, nextPosQueue)
						uniquePosesToMove[nextPosQueue] = struct{}{}
					}

					nextPosBox := [2]int{item[0], item[1] + 1}
					if _, ok := uniquePosesToMove[nextPosBox]; !ok {
						queue = append(queue, nextPosBox)
						uniquePosesToMove[nextPosBox] = struct{}{}
					}
				case "]":
					posesToMove = append(posesToMove, item)

					nextPosQueue := [2]int{item[0] + directions[v][0], item[1] + directions[v][1]}
					nextPoses[item] = nextPosQueue

					if _, ok := uniquePosesToMove[nextPosQueue]; !ok {
						queue = append(queue, nextPosQueue)
						uniquePosesToMove[nextPosQueue] = struct{}{}
					}

					nextPosBox := [2]int{item[0], item[1] - 1}
					if _, ok := uniquePosesToMove[nextPosBox]; !ok {
						queue = append(queue, nextPosBox)
						uniquePosesToMove[nextPosBox] = struct{}{}
					}
				case "#":
					posesToMove = [][2]int{}
					nextPoses = map[[2]int][2]int{}
					uniquePosesToMove = map[[2]int]struct{}{}

					done = true
				case ".":
					uniquePosesToMove[item] = struct{}{}
				case "@":
					posesToMove = append(posesToMove, item)
					nextPosQueue := [2]int{item[0] + directions[v][0], item[1] + directions[v][1]}
					nextPoses[item] = nextPosQueue
					queue = append(queue, nextPosQueue)

					uniquePosesToMove[item] = struct{}{}
				}

				if done {
					break
				}
			}

			previousPoses := map[[2]int]string{}
			newPoses := map[[2]int]string{}
			allPoseChanges := [][2]int{}
			for currentPos, nextPos := range nextPoses {
				previousPoses[currentPos] = "."
				newPoses[nextPos] = matrix[currentPos[0]][currentPos[1]]
				allPoseChanges = append(allPoseChanges, currentPos)
				allPoseChanges = append(allPoseChanges, nextPos)
			}

			//matrixStr = sprintMatrix(matrix, map[string][][2]int{Blue: allPoseChanges})
			//fmt.Printf("Move %v %v Before:\n%s\n\n", j, v, matrixStr)

			for pos, newV := range previousPoses {
				matrix[pos[0]][pos[1]] = newV
			}
			for pos, newV := range newPoses {
				matrix[pos[0]][pos[1]] = newV
				if newV == "@" {
					robotPos = pos
				}
			}

			//matrixStr = sprintMatrix(matrix, map[string][][2]int{Green: allPoseChanges})
			//fmt.Printf("Move %v %v After:\n%s\n", j, v, matrixStr)
			//fmt.Printf("\n")
		}
	}

	sum := 0
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] != "[" {
				continue
			}
			sum += (100 * i) + j
		}
	}

	//matrixStr = sprintMatrix(matrix)
	//fmt.Printf("Result:\n%s\n", matrixStr)

	fmt.Println(sum)

	return nil
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func sprintMatrix(matrix [][]string, highlights ...map[string][][2]int) string {
	readyHighlights := map[[2]int]string{}
	for _, highlight := range highlights {
		for hColor, coords := range highlight {
			for _, coord := range coords {
				readyHighlights[coord] = hColor
			}
		}
	}

	results := []string{}
	for i := range matrix {
		result := ""
		for j := range matrix[i] {
			cColor := ""
			if highlight, exists := readyHighlights[[2]int{i, j}]; exists {
				cColor = highlight
				result = result + cColor
			}

			result = result + matrix[i][j]

			if cColor != "" {
				result = result + Reset
				cColor = Reset
			}
		}
		results = append(results, result)
	}

	return strings.Join(results, "\n")
}
