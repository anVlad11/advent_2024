package main

import (
	"fmt"
	"github.com/anVlad11/advent_2024/pkg/utils"
	"math"
	"strings"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func do() error {
	input, err := utils.GetInput("inputs/day_02/input_1.txt")
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
	safeCount := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}

		lineItems := utils.ConvertSlice(strings.Split(line, " "), utils.MustParseInt64)

		if safeCheck(lineItems) {
			safeCount++
		}

	}

	fmt.Println(safeCount)

	return nil
}

func part2(lines []string) error {
	safeCount := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}

		//fmt.Println(line)

		lineItems := utils.ConvertSlice(strings.Split(line, " "), utils.MustParseInt64)

		if safeCheck(lineItems) {
			//fmt.Println("safe by itself")
			safeCount++
			continue
		}

		for i := range lineItems {
			lineItemCopy := make([]int64, len(lineItems))
			copy(lineItemCopy, lineItems)
			optionalLineItem := append(lineItemCopy[:i], lineItemCopy[i+1:]...)
			//fmt.Printf("optional %v ", optionalLineItem)
			if safeCheck(optionalLineItem) {
				safeCount++
				//fmt.Printf("safe\n")
				break
			}
			//fmt.Printf("not safe\n")
		}

	}

	fmt.Println(safeCount)

	return nil
}

func safeCheck(lineItems []int64) bool {
	safe := true

	direction := ""
	for i := 1; i < len(lineItems); i++ {
		newDirection := lineItems[i] > lineItems[i-1]
		if direction == "" {
			if newDirection {
				direction = "up"
			} else {
				direction = "down"
			}
		}

		if (direction == "up" && !newDirection) || (direction == "down" && newDirection) {
			safe = false
			break
		}

		distance := int64(math.Abs(float64(lineItems[i] - lineItems[i-1])))
		if distance < 1 || distance > 3 {
			safe = false
			break
		}
	}

	return safe
}
