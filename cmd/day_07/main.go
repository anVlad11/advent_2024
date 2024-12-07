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
	input, err := utils.GetInput("inputs/day_07/input_1.txt")
	if err != nil {
		return err
	}

	err = part1(input, operatorsDay1)
	if err != nil {
		return err
	}

	err = part1(input, operatorsDay2)
	if err != nil {
		return err
	}

	return nil
}

type Operator string

const (
	OperatorSum     Operator = "+"
	OperatorProduct Operator = "*"
	OperatorConcat  Operator = "||"
)

var operatorsDay1 = []Operator{
	OperatorSum,
	OperatorProduct,
}

var operatorsDay2 = []Operator{
	OperatorSum,
	OperatorProduct,
	OperatorConcat,
}

var operations = map[Operator]func(operands ...int64) int64{
	OperatorSum: func(operands ...int64) int64 {
		sum := int64(0)
		for _, operand := range operands {
			sum += operand
		}

		return sum
	},
	OperatorProduct: func(operands ...int64) int64 {
		if len(operands) == 0 {
			return 0
		}
		product := int64(1)
		for _, operand := range operands {
			product *= operand
		}

		return product
	},
	OperatorConcat: func(operands ...int64) int64 {
		if len(operands) == 0 {
			return 0
		}
		concat := ""
		for _, operand := range operands {
			concat = concat + fmt.Sprintf("%d", operand)
		}

		return utils.MustParseInt64(concat)
	},
}

func part1(lines []string, dailyOperators []Operator) error {
	sum := int64(0)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		halves := strings.Split(line, ":")
		operands := utils.ConvertSlice(strings.Split(strings.TrimSpace(halves[1]), " "), utils.MustParseInt64)

		dailyOperations := map[Operator]func(operands ...int64) int64{}
		for _, operator := range dailyOperators {
			dailyOperations[operator] = operations[operator]
		}

		tree := &Tree{
			Root:       nil,
			Target:     utils.MustParseInt64(halves[0]),
			Operations: dailyOperations,
		}

		root := &Node{
			Value:      operands[0],
			RightLinks: make([]*Link, 0, len(tree.Operations)),
		}

		tree.Root = root

		parentNodes := []*Node{tree.Root}

		for i := 1; i < len(operands); i++ {
			newParentNodes := make([]*Node, 0, len(parentNodes)*len(tree.Operations))
			for _, parentNode := range parentNodes {
				for operator := range tree.Operations {
					newNode := &Node{
						Value:      operands[i],
						RightLinks: make([]*Link, 0, len(tree.Operations)),
					}

					newLink := &Link{
						LeftNode:  parentNode,
						Operator:  operator,
						RightNode: newNode,
					}

					parentNode.RightLinks = append(parentNode.RightLinks, newLink)
					newNode.LeftLink = newLink

					newParentNodes = append(newParentNodes, newNode)
				}
			}
			parentNodes = newParentNodes
		}

		solution := tree.SolutionNode()
		//fmt.Printf("%s", line)
		if solution != nil {
			sum += tree.Target
			//fmt.Printf(" is solvable by: %s", solution.GetPath())
		}
		//fmt.Printf("\n")
	}

	fmt.Println(sum)

	return nil
}

func (m *Node) GetPath() string {
	path := fmt.Sprintf("%d", m.Value)

	if m.LeftLink != nil {
		path = fmt.Sprintf("%s%s%s", m.LeftLink.LeftNode.GetPath(), m.LeftLink.Operator, path)
	}

	return path
}

func part2(lines []string) error {
	return nil
}

type Node struct {
	Value      int64
	LeftLink   *Link
	RightLinks []*Link
}

type Link struct {
	LeftNode  *Node
	Operator  Operator
	RightNode *Node
	Value     int64
}

type Tree struct {
	Root       *Node
	Target     int64
	Operations map[Operator]func(operands ...int64) int64
}

func (m *Tree) SolutionNode() *Node {
	for _, link := range m.Root.RightLinks {
		leftValue := m.Root.Value
		if m.Root.LeftLink != nil {
			leftValue = m.Root.LeftLink.Value
		}

		link.Value = m.Operations[link.Operator](leftValue, link.RightNode.Value)

		/*
			fmt.Printf("%s\t", link.RightNode.GetPath())
			fmt.Printf("[ %s = %d ] %s %d = %d",
				link.LeftNode.GetPath(),
				leftValue,
				link.Operator,
				link.RightNode.Value,
				link.Value,
			)

		*/

		if link.Value > m.Target {
			//fmt.Printf(" - too much\n\n")
			continue
		}

		if (link.Value == m.Target) && len(link.RightNode.RightLinks) == 0 {
			//fmt.Printf(" - is a solution!\n")
			return link.RightNode
		}
		//fmt.Printf(" - go deeper\n")

		subTree := &Tree{
			Root:       link.RightNode,
			Target:     m.Target,
			Operations: m.Operations,
		}

		subTreeSolution := subTree.SolutionNode()
		if subTreeSolution != nil {
			return subTreeSolution
		}

	}

	return nil
}
