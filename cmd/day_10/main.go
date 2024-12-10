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
	input, err := utils.GetInput("inputs/day_10/input_1.txt")
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

var validDirections = [][2]int{{0, -1}, {1, 0}, {-1, 0}, {0, 1}}

func part1(lines []string) error {
	matrix := make([][]int, len(lines))

	starts := map[[2]int]int{}

	for i, line := range lines {
		matrix[i] = make([]int, len(line))
		for j, point := range strings.Split(line, "") {
			matrix[i][j] = utils.MustParseInt(point)
			if matrix[i][j] == 0 {
				starts[[2]int{i, j}] = 0
			}
		}
	}

	sum := 0
	for startCoords := range starts {
		visited := map[[2]int][2]int{}
		queue := [][2]int{startCoords}
		previousNode := [2]int{-1, -1}

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			if _, exists := visited[node]; exists {
				continue
			}

			if previousNode != [2]int{-1, -1} {
				visited[node] = previousNode
			}
			previousNode = node

			if matrix[node[0]][node[1]] == 9 {
				sum++
				continue
			}

			for _, direction := range validDirections {
				newNode := [2]int{node[0] + direction[0], node[1] + direction[1]}

				if newNode[0] >= len(matrix) || newNode[0] < 0 || newNode[1] >= len(matrix) || newNode[1] < 0 {
					continue
				}
				if _, exists := visited[newNode]; exists {
					continue
				}
				if matrix[newNode[0]][newNode[1]]-matrix[node[0]][node[1]] != 1 {
					continue
				}
				queue = append(queue, newNode)
			}
		}
	}

	fmt.Println(sum)

	return nil
}

func part2(lines []string) error {
	matrix := make([][]int, len(lines))

	starts := map[[2]int]int{}

	for i, line := range lines {
		matrix[i] = make([]int, len(line))
		for j, point := range strings.Split(line, "") {
			matrix[i][j] = utils.MustParseInt(point)
			if matrix[i][j] == 0 {
				starts[[2]int{i, j}] = 0
			}
		}
	}

	sum := 0
	for startCoords := range starts {
		queue := [][2]int{startCoords}

		thread := 1
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			if matrix[node[0]][node[1]] == 9 {
				continue
			}

			newNodes := 0
			for _, direction := range validDirections {
				newNode := [2]int{node[0] + direction[0], node[1] + direction[1]}

				if newNode[0] >= len(matrix) || newNode[0] < 0 || newNode[1] >= len(matrix) || newNode[1] < 0 {
					continue
				}
				if matrix[newNode[0]][newNode[1]]-matrix[node[0]][node[1]] != 1 {
					continue
				}
				queue = append(queue, newNode)
				newNodes++
			}
			thread += newNodes - 1
		}
		sum += thread
	}

	fmt.Println(sum)

	return nil
}
