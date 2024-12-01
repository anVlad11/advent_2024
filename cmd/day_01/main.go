package main

import (
	"fmt"
	"github.com/anVlad11/advent_2024/pkg/data"
	"io"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func do() error {
	err := part1()
	if err != nil {
		return err
	}

	err = part2()
	if err != nil {
		return err
	}

	return nil
}

func part1() error {
	inputs, err := data.Inputs.Open("inputs/day_01/input_1.txt")
	if err != nil {
		return err
	}

	input, err := io.ReadAll(inputs)
	if err != nil {
		return err
	}

	inputString := string(input)
	lines := strings.Split(inputString, "\n")
	left := make([]int64, len(lines))
	right := make([]int64, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}

		lineItems := strings.Split(line, "   ")
		left[i], err = strconv.ParseInt(lineItems[0], 10, 64)
		if err != nil {
			return err
		}
		right[i], err = strconv.ParseInt(lineItems[1], 10, 64)
		if err != nil {
			return err
		}

	}

	slices.Sort(left)
	slices.Sort(right)
	distance := 0
	for i := range left {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Println(distance)

	return nil
}

func part2() error {
	inputs, err := data.Inputs.Open("inputs/day_01/input_1.txt")
	if err != nil {
		return err
	}

	input, err := io.ReadAll(inputs)
	if err != nil {
		return err
	}

	inputString := string(input)
	lines := strings.Split(inputString, "\n")
	leftScoreboard := map[int64]int64{}
	similarityScoreboard := map[int64]int64{}
	rightScoreboard := map[int64]int64{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}

		lineItems := strings.Split(line, "   ")
		var left int64
		var right int64

		left, err = strconv.ParseInt(lineItems[0], 10, 64)
		if err != nil {
			return err
		}
		right, err = strconv.ParseInt(lineItems[1], 10, 64)
		if err != nil {
			return err
		}

		if _, exists := leftScoreboard[left]; !exists {
			leftScoreboard[left] = 1
		} else {
			leftScoreboard[left]++
		}

		if _, exists := similarityScoreboard[left]; !exists {
			similarityScoreboard[left] = 0
		}

		if _, exists := rightScoreboard[right]; !exists {
			rightScoreboard[right] = 1
		} else {
			rightScoreboard[right]++
		}

		if _, exists := rightScoreboard[left]; exists {
			similarityScoreboard[left] = rightScoreboard[left]
		}

		if _, exists := similarityScoreboard[right]; exists {
			similarityScoreboard[right] = rightScoreboard[right]
		}
	}

	score := int64(0)
	for k, v := range leftScoreboard {
		score += (k * similarityScoreboard[k]) * v
	}

	fmt.Println(score)

	return nil
}
