package main

import (
	"fmt"
	"github.com/anVlad11/advent_2024/pkg/utils"
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
	input, err := utils.GetInput("inputs/day_11/input_1.txt")
	if err != nil {
		return err
	}

	err = part1(input, 25)
	if err != nil {
		return err
	}

	err = part2(input, 75)
	if err != nil {
		return err
	}

	return nil
}

func part1(lines []string, iterations int64) error {
	line := strings.Split(lines[0], " ")

	strOutput := make([]string, len(line))
	output := make([]int64, len(line))
	tracker := make([]int64, len(line))

	for i, v := range line {
		strOutput[i] = v
		output[i] = utils.MustParseInt64(v)
		tracker[i] = iterations
	}

	for i := 0; i < len(output); {
		if tracker[i] <= 0 {
			i++
			continue
		}

		v := output[i]
		if v == 0 {
			output[i] = 1
			strOutput[i] = strconv.FormatInt(output[i], 10)
			tracker[i]--
		} else if len(strOutput[i])%2 == 0 {
			half := len(strOutput[i]) / 2
			newOutput := []int64{utils.MustParseInt64(strOutput[i][:half]), utils.MustParseInt64(strOutput[i][half:])}
			strOutput = slices.Replace(strOutput, i, i+1, strconv.FormatInt(newOutput[0], 10), strconv.FormatInt(newOutput[1], 10))
			output = slices.Replace(output, i, i+1, newOutput[0], newOutput[1])
			tracker = slices.Replace(tracker, i, i+1, tracker[i]-1, tracker[i]-1)
		} else {
			output[i] = output[i] * 2024
			strOutput[i] = strconv.FormatInt(output[i], 10)
			tracker[i]--
		}
	}

	fmt.Println(len(output))

	return nil
}

func part2(lines []string, iterations int64) error {
	line := utils.ConvertSlice(strings.Split(lines[0], " "), utils.MustParseInt64)

	numCount := make(map[int64]int64)
	for _, v := range line {
		numCount[v]++
	}

	for i := int64(0); i < iterations; i++ {
		newNumCount := make(map[int64]int64)
		for num, count := range numCount {
			if num == 0 {
				newNumCount[1] += count
			} else if len(utils.FormatInt64(num))%2 == 0 {
				strNum := utils.FormatInt64(num)
				half := len(strNum) / 2

				newNumCount[utils.MustParseInt64(strNum[:half])] += count
				newNumCount[utils.MustParseInt64(strNum[half:])] += count
			} else {
				newNumCount[num*2024] += count
			}
		}

		numCount = newNumCount
	}

	sum := int64(0)
	for _, count := range numCount {
		sum += count
	}

	fmt.Println(sum)

	return nil
}
