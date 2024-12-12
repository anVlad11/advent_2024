package main

import (
	"fmt"
	"github.com/anVlad11/advent_2024/pkg/utils"
	"sort"
	"strings"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func do() error {
	input, err := utils.GetInput("inputs/day_12/input_1.txt")
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

var directions = [][2]int{{0, -1}, {1, 0}, {-1, 0}, {0, 1}}

func part1(lines []string) error {
	visited := map[[2]int]map[int]int{}
	regions := [][2]int{}
	regionLetters := []string{}
	region := -1

	//count := 0

	for i, line := range lines {
		for j := range line {
			if _, exists := visited[[2]int{i, j}]; exists {
				continue
			}

			queue := [][2]int{{i, j}}
			region++

			for len(queue) > 0 {
				item := queue[0]
				queue = queue[1:]

				regionLetter := string(lines[item[0]][item[1]])
				if len(regions) < region+1 {
					regions = append(regions, [2]int{0, 0})
					regionLetters = append(regionLetters, regionLetter)
				}

				regions[region][0]++

				if _, exists := visited[item]; !exists {
					visited[item] = map[int]int{}
				}
				visited[item][region]++
				regionString := string(lines[item[0]][item[1]])

				highlights := map[[2]int]string{}

				for coord, _ := range visited {
					highlights[coord] = Gray
				}

				highlights[item] = Green

				for _, direction := range directions {
					newI := item[0] + direction[0]
					newJ := item[1] + direction[1]

					neighbourCoord := [2]int{newI, newJ}
					highlights[neighbourCoord] = Red

					if newI < 0 || newI >= len(lines) || newJ < 0 || newJ >= len(lines) {
						regions[region][1]++
						continue
					}

					if string(lines[newI][newJ]) != regionString {
						regions[region][1]++
					} else {
						if _, exists := visited[neighbourCoord][region]; !exists {
							queue = append(queue, neighbourCoord)
							visited[neighbourCoord] = map[int]int{region: 0}
						}
					}
				}

				/*
					if count > 1000 {
						count = 0
						fmt.Printf("%v\n%v\n%s\n", queue, regions[region], sprintfMap(lines, highlights))
					}
					count++
				*/
			}
		}
	}

	//fmt.Println(regions)

	sum := 0
	for i := range regions {
		//fmt.Printf("%s %v * %v = %v\n", regionLetters[i], regions[i][0], regions[i][1], regions[i][0]*regions[i][1])
		sum += regions[i][0] * regions[i][1]
	}

	fmt.Println(sum)

	return nil
}

func part2(lines []string) error {
	visited := map[[2]int]map[int]int{}
	regions := [][2]int{}
	regionLetters := []string{}
	region := -1

	//count := 0

	for i, line := range lines {
		for j := range line {
			if _, exists := visited[[2]int{i, j}]; exists {
				continue
			}

			queue := [][2]int{{i, j}}
			region++

			regionBorders := map[[2]float64][][2]int{}

			for len(queue) > 0 {
				item := queue[0]
				queue = queue[1:]

				regionLetter := string(lines[item[0]][item[1]])
				if len(regions) < region+1 {
					regions = append(regions, [2]int{0, 0})
					regionLetters = append(regionLetters, regionLetter)
				}

				regions[region][0]++

				if _, exists := visited[item]; !exists {
					visited[item] = map[int]int{}
				}
				visited[item][region]++
				regionString := string(lines[item[0]][item[1]])

				highlights := map[[2]int]string{}

				for coord, _ := range visited {
					highlights[coord] = Gray
				}

				highlights[item] = Green

				for _, direction := range directions {
					newI := item[0] + direction[0]
					newJ := item[1] + direction[1]

					neighbourCoord := [2]int{newI, newJ}
					highlights[neighbourCoord] = Red

					if newI < 0 || newI >= len(lines) || newJ < 0 || newJ >= len(lines) {
						kI := float64(direction[0]) * 0.2
						kJ := float64(direction[1]) * 0.2
						borderPseudoCoord := [2]float64{0, 0}
						if kI != 0 {
							borderPseudoCoord[0] = float64(neighbourCoord[0]) + kI
						}
						if kJ != 0 {
							borderPseudoCoord[1] = float64(neighbourCoord[1]) + kJ
						}

						regionBorders[borderPseudoCoord] = append(regionBorders[borderPseudoCoord], neighbourCoord)

						continue
					}

					if string(lines[newI][newJ]) != regionString {
						kI := float64(direction[0]) * 0.2
						kJ := float64(direction[1]) * 0.2
						borderPseudoCoord := [2]float64{0, 0}
						if kI != 0 {
							borderPseudoCoord[0] = float64(neighbourCoord[0]) + kI
						}
						if kJ != 0 {
							borderPseudoCoord[1] = float64(neighbourCoord[1]) + kJ
						}

						regionBorders[borderPseudoCoord] = append(regionBorders[borderPseudoCoord], neighbourCoord)
					} else {
						if _, exists := visited[neighbourCoord][region]; !exists {
							queue = append(queue, neighbourCoord)
							visited[neighbourCoord] = map[int]int{region: 0}
						}
					}
				}

				/*
					if count > 1000 {
						count = 0
						fmt.Printf("%v\n%v\n%s\n", queue, regions[region], sprintfMap(lines, highlights))
					}
					count++
				*/
			}

			for pseudoCoord, coords := range regionBorders {
				sortCoords := make([]int, len(coords))

				idx := 0
				if pseudoCoord[0] != 0 {
					idx = 1
				}

				for k, coord := range coords {
					sortCoords[k] = coord[idx]
				}

				sort.Ints(sortCoords)
				borderCount := 1
				for m := 1; m < len(sortCoords); m++ {
					mCur := sortCoords[m]
					mPrev := sortCoords[m-1]
					mDiff := mPrev - mCur
					if !((mDiff == 1) || (mDiff == -1)) {
						borderCount++
					}
				}
				regions[region][1] += borderCount
			}
		}
	}

	//fmt.Println(regions)

	sum := 0
	for i := range regions {
		//fmt.Printf("%s %v * %v = %v\n", regionLetters[i], regions[i][0], regions[i][1], regions[i][0]*regions[i][1])
		sum += regions[i][0] * regions[i][1]
	}

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

func sprintfMap(matrix []string, highlights ...map[[2]int]string) string {
	results := []string{}
	color := ""
	for i, line := range matrix {
		for j, v := range line {
			result := ""
			for _, highlight := range highlights {
				if c, exists := highlight[[2]int{i, j}]; exists {
					result += c
					color = c
				}
			}

			result += string(v)

			if color != "" {
				result += Reset
				color = ""
			}
			results = append(results, result)
		}
		results = append(results, "\n")
	}

	result := strings.Join(results, "")

	return result
}
