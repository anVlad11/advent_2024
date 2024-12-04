package main

import (
	"fmt"
	"github.com/anVlad11/advent_2024/pkg/utils"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func do() error {
	input, err := utils.GetInput("inputs/day_04/input_1.txt")
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

type Masks [][][2]int

func generateMasks(directions []int, length int, alignToCenter bool) [][][2]int {
	masks := [][][2]int{}

	for _, directionI := range directions {
		for _, directionJ := range directions {
			mask := make([][2]int, length)
			for pos := range mask {
				origin := [2]int{0, 0}
				if alignToCenter {
					origin = [2]int{-directionI * (length / 2), -directionJ * (length / 2)}
				}
				mask[pos] = [2]int{origin[0] + directionI*pos, origin[1] + directionJ*pos}
			}
			masks = append(masks, mask)
		}
	}

	return masks
}

func checkMasks(matrix [][]string, masks Masks, origin [2]int, word string) int {
	result := 0
	for _, mask := range masks {
		wordResult := 0
		for i, maskLetter := range mask {
			positionI := maskLetter[0] + origin[0]
			positionJ := maskLetter[1] + origin[1]

			if positionI < 0 || positionI >= len(matrix) {
				continue
			}

			if positionJ < 0 || positionJ >= len(matrix[positionI]) {
				continue
			}

			letter := string(word[i])
			if matrix[positionI][positionJ] == letter {
				wordResult++
			}
		}
		if wordResult == len(word) {
			result++
		}
	}

	return result
}

func part1(lines []string) error {
	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = make([]string, len(lines[i]))
		for j := range line {
			matrix[i][j] = string(line[j])
		}
	}

	result := 0
	masks := generateMasks([]int{-1, 0, 1}, 4, false)
	for i := range matrix {
		for j := range matrix[i] {
			positionResult := checkMasks(matrix, masks, [2]int{i, j}, "XMAS")
			result += positionResult
		}
	}

	fmt.Println(result)

	return nil
}

func part2(lines []string) error {
	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = make([]string, len(lines[i]))
		for j := range line {
			matrix[i][j] = string(line[j])
		}
	}

	result := 0
	masks := generateMasks([]int{-1, 1}, 3, true)
	for i := range matrix {
		for j := range matrix[i] {
			positionResult := checkMasks(matrix, masks, [2]int{i, j}, "MAS")
			if positionResult >= 2 {
				result++
			}
		}
	}

	fmt.Println(result)

	return nil
}
