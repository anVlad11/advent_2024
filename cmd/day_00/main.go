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
	input, err := utils.GetInput("inputs/day_00/input_1.txt")
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

	return nil
}

func part2(lines []string) error {

	return nil
}
