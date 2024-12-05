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
	input, err := utils.GetInput("inputs/day_05/input_1.txt")
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

type Rule struct {
	Left  int64
	Right int64
}

type Ruleset struct {
	LeftRules  map[int64][]*Rule
	RightRules map[int64][]*Rule
}

func part1(lines []string) error {
	ruleset := &Ruleset{
		LeftRules:  make(map[int64][]*Rule),
		RightRules: make(map[int64][]*Rule),
	}

	i := 0
	for ; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			i++
			break
		}
		ruleRaw := utils.ConvertSlice(strings.Split(line, "|"), utils.MustParseInt64)
		rule := &Rule{
			Left:  ruleRaw[0],
			Right: ruleRaw[1],
		}
		ruleset.LeftRules[rule.Left] = append(ruleset.LeftRules[rule.Left], rule)
		ruleset.RightRules[rule.Right] = append(ruleset.RightRules[rule.Right], rule)
	}

	validSum := int64(0)
	for ; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			break
		}
		elements := utils.ConvertSlice(strings.Split(line, ","), utils.MustParseInt64)
		valid := true
		for j, element := range elements {
			for _, rule := range ruleset.LeftRules[element] {
				for k := 0; k < j; k++ {
					if rule.Right == elements[k] {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			validSum += elements[len(elements)/2]
		}
	}

	fmt.Println(validSum)

	return nil
}

func part2(lines []string) error {
	ruleset := &Ruleset{
		LeftRules:  make(map[int64][]*Rule),
		RightRules: make(map[int64][]*Rule),
	}

	i := 0
	for ; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			i++
			break
		}
		ruleRaw := utils.ConvertSlice(strings.Split(line, "|"), utils.MustParseInt64)
		rule := &Rule{
			Left:  ruleRaw[0],
			Right: ruleRaw[1],
		}
		ruleset.LeftRules[rule.Left] = append(ruleset.LeftRules[rule.Left], rule)
		ruleset.RightRules[rule.Right] = append(ruleset.RightRules[rule.Right], rule)
	}

	validSum := int64(0)
	for ; i < len(lines); i++ {
		line := lines[i]
		if len(line) == 0 {
			break
		}
		elements := utils.ConvertSlice(strings.Split(line, ","), utils.MustParseInt64)
		valid := true
		for j, element := range elements {
			for _, rule := range ruleset.LeftRules[element] {
				for k := 0; k < j; k++ {
					if rule.Right == elements[k] {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
			}
			if !valid {
				break
			}
		}

		if !valid {
			sort.Slice(elements, func(i, j int) bool {
				for _, rule := range ruleset.LeftRules[elements[i]] {
					if rule.Right == elements[j] {
						return false
					}

				}
				return true
			})

			validSum += elements[len(elements)/2]
		}
	}

	fmt.Println(validSum)

	return nil
}
