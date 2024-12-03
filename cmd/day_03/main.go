package main

import (
	"fmt"
	"github.com/anVlad11/advent_2024/pkg/utils"
	"regexp"
	"strings"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func do() error {
	input, err := utils.GetInput("inputs/day_03/input_1.txt")
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
	regex := regexp.MustCompile(`mul\((\d*),(\d*)\)`)

	line := strings.Join(lines, "")

	sum := int64(0)
	multiplications := regex.FindAllStringSubmatch(line, -1)

	for _, multiplication := range multiplications {
		left := utils.MustParseInt64(multiplication[1])
		right := utils.MustParseInt64(multiplication[2])

		sum += left * right
	}

	fmt.Println(sum)

	return nil
}

func part2(lines []string) error {
	regex := regexp.MustCompile(`((mul)\((\d*),(\d*)\))|((don't)\(\))|((do)\(\))`)

	line := strings.Join(lines, "")

	sum := int64(0)
	operations := regex.FindAllStringSubmatch(line, -1)

	enabled := true

	for _, operation := range operations {
		switch {
		case operation[2] == "mul":
			if !enabled {
				continue
			}
			left := utils.MustParseInt64(operation[3])
			right := utils.MustParseInt64(operation[4])

			sum += left * right
		case operation[8] == "do":
			enabled = true
		case operation[6] == "don't":
			enabled = false
		default:
			fmt.Printf("Invalid operation: %v\n", operation)
		}
	}

	fmt.Println(sum)

	return nil
}
