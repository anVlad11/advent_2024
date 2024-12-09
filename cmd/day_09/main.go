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
	input, err := utils.GetInput("inputs/day_09/input_1.txt")
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

func part1(lines []string) error {
	line := lines[0]

	disk := []int{}

	for i, vRaw := range line {
		v := utils.MustParseInt64(string(vRaw))

		if i%2 != 0 {
			for j := int64(0); j < v; j++ {
				disk = append(disk, -1)
			}
			continue
		}

		for j := int64(0); j < v; j++ {
			disk = append(disk, i/2)
		}
	}

	i := 0
	j := len(disk) - 1
	sum := int64(0)
	for {
		if i >= len(disk) {
			break
		}
		if disk[i] != -1 {
			sum += int64(disk[i] * i)
			i++
			continue
		}
		if j < 0 {
			break
		}
		if disk[j] == -1 {
			j--
			continue
		}
		if i >= j {
			break
		}
		disk[i] = disk[j]
		sum += int64(disk[i] * i)
		disk[j] = -1

		//fmt.Printf("%d %d %d %v\n", i, disk[i]*i, sum, disk)

		i++
		j--

	}

	fmt.Println(sum)

	return nil
}

func part2(lines []string) error {
	line := lines[0]

	disk := []int{}
	denseDisk := make([]int, len(line))
	diskIDOffsets := make([]int, len(line))

	for i, vRaw := range line {
		v := utils.MustParseInt(string(vRaw))
		denseDisk[i] = v
		diskIDOffsets[i] = len(disk)

		if i%2 != 0 {
			for j := 0; j < v; j++ {
				disk = append(disk, -1)
			}
			continue
		}

		for j := 0; j < v; j++ {
			disk = append(disk, i/2)
		}
	}

	//fmt.Printf("%s\n\n", sprintfDisk(disk))

	sum := int64(0)
	for j := len(disk) - 1; j >= 0; {
		//fmt.Printf("%s - %d\n\n", sprintfDisk(disk, map[int]string{j: Blue}), j)
		if disk[j] == -1 {
			j--
			continue
		}

		// Global ID - position of the block in dense disk
		// Local ID - position of the block in dense disk, by type

		sourceBlockGlobalID := disk[j] * 2
		sourceBlockSize := denseDisk[sourceBlockGlobalID]
		targetBlockSize := 0
		for globalBlockID := 1; globalBlockID < len(denseDisk); globalBlockID += 2 {
			targetBlockSize = denseDisk[globalBlockID]
			if targetBlockSize < sourceBlockSize {
				continue
			}

			targetBlockDiskOffset := diskIDOffsets[globalBlockID]
			if targetBlockDiskOffset > j {
				targetBlockSize = 0
				break
			}

			mem := make([]int, sourceBlockSize)
			copy(mem, disk[targetBlockDiskOffset:targetBlockDiskOffset+sourceBlockSize])
			copy(disk[targetBlockDiskOffset:targetBlockDiskOffset+sourceBlockSize], disk[j-sourceBlockSize+1:j+1])
			copy(disk[j-sourceBlockSize+1:j+1], mem)

			/*
				// example:
				// (4*7)+(5*7)+(6*7)+(7*7)+(8*7) is generalized into
				// (7/2) * (8-4+1) * (4+8)
				minJ := targetBlockDiskOffset
				maxJ := targetBlockDiskOffset + sourceBlockSize - 1
				sumPart := (float64(disk[targetBlockDiskOffset]) / 2.0) * //(7/2)
					float64(maxJ-minJ+1) * //(8-4+1)
					float64(minJ+maxJ) //(8+4)

				fmt.Printf("%s\n", sprintfDisk(disk, map[int]string{minJ: Green, maxJ: Green}))
				fmt.Printf("Calc: (%v/2) * (%v-%v+1) * (%v+%v) = %v\n",
					disk[targetBlockDiskOffset],
					maxJ,
					minJ,
					maxJ,
					minJ,
					sumPart,
				)

				fmt.Printf("%v + %v = %v\n", sum, sumPart, sum)
				sum += int64(sumPart)
				fmt.Printf("\n")

			*/
			/*
				for k := targetBlockDiskOffset; k < targetBlockDiskOffset+sourceBlockSize; k++ {
					//fmt.Printf("%s\n", sprintfDisk(disk, map[int]string{k: Yellow}))

					sumPart := int64(disk[k] * k)
					//fmt.Printf("Moved: %v * %v = %v\n", k, disk[k], sumPart)
					//fmt.Printf("%v + %v = %v\n\n", sum, sumPart, sum)
					sum += sumPart
				}

			*/

			/*
				fmt.Printf(
					"found appropriate empty block: %d -> %d, %v %v\n",
					sourceBlockGlobalID,
					globalBlockID,
					denseDisk,
					disk,
				)

			*/

			denseDisk[globalBlockID] -= sourceBlockSize
			diskIDOffsets[globalBlockID] += sourceBlockSize
			j -= sourceBlockSize
			break
		}

		if targetBlockSize < sourceBlockSize {
			v := j
			for {
				if v < 0 {
					break
				}
				if disk[v] != disk[j] {
					break
				}
				sumPart := int64(disk[v] * v)
				//fmt.Println(disk)
				//fmt.Printf("%s\n", sprintfDisk(disk, map[int]string{v: Red}))
				//fmt.Printf("Stayed: %v * %v = %v\n", disk[v], v, sumPart)
				//fmt.Printf("%v + %v = %v\n", sum, sumPart, sum+sumPart)
				//fmt.Printf("\n")
				sum += sumPart
				v--
			}
			j = v
		}
	}

	// :DDDDDD
	sum = int64(0)

	for i, v := range disk {
		if v > -1 {
			sum += int64(i * v)
		}
	}

	//fmt.Println(disk)
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

func sprintfDisk(disk []int, highlights ...map[int]string) string {
	results := make([]string, len(disk))
	color := ""
	for i, v := range disk {
		result := ""
		for _, highlight := range highlights {
			if c, exists := highlight[i]; exists {
				result += c
				color = c
			}
		}

		if v == -1 {
			result += "."
		} else {
			result += fmt.Sprintf("%v", v)
		}
		if color != "" {
			result += Reset
			color = ""
		}
		results[i] = result
	}

	result := strings.Join(results, " ")

	return result
}
