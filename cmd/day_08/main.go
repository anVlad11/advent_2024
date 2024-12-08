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
	input, err := utils.GetInput("inputs/day_08/input_1.txt")
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
	matrix := make([][]string, len(lines))

	emitters := map[string][][2]int{}
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		matrix[i] = strings.Split(line, "")
		for j, v := range matrix[i] {
			if v != "." {
				emitters[v] = append(emitters[v], [2]int{i, j})
			}
		}
	}

	antinodes := map[[2]int]struct{}{}
	for _, letterEmitters := range emitters {
		for i := 0; i < len(letterEmitters)-1; i++ {
			iCoords := letterEmitters[i]
			xa := float64(iCoords[0])
			ya := float64(iCoords[1])
			for j := i + 1; j < len(letterEmitters); j++ {
				jCoords := letterEmitters[j]
				xb := float64(jCoords[0])
				yb := float64(jCoords[1])

				distX := math.Abs(xb - xa)
				distY := math.Abs(yb - ya)

				xcFloat := xa - distX
				if xa > xb {
					xcFloat = xa + distX
				}
				ycFloat := ya - distY
				if ya > yb {
					ycFloat = ya + distY
				}

				xc := int(xcFloat)
				yc := int(ycFloat)

				if xc < len(lines) && xc >= 0 && yc < len(lines) && yc >= 0 {
					antinodes[[2]int{xc, yc}] = struct{}{}
				}

				xdFloat := xb + distX
				if xb < xa {
					xdFloat = xb - distX
				}
				ydFloat := yb + distY
				if yb < ya {
					ydFloat = yb - distY
				}

				xd := int(xdFloat)
				yd := int(ydFloat)

				if xd < len(lines) && xd >= 0 && yd < len(lines) && yd >= 0 {
					antinodes[[2]int{xd, yd}] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(antinodes))

	return nil
}

func part2(lines []string) error {
	matrix := make([][]string, len(lines))

	emitters := map[string][][2]int{}
	antinodes := map[[2]int]struct{}{}
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		matrix[i] = strings.Split(line, "")
		//fmt.Printf("%v\n", matrix[i])
		for j, v := range matrix[i] {
			if v != "." {
				emitters[v] = append(emitters[v], [2]int{i, j})
				antinodes[[2]int{i, j}] = struct{}{}
			}
		}
	}

	for _, letterEmitters := range emitters {
		for i := 0; i < len(letterEmitters)-1; i++ {
			iCoords := letterEmitters[i]
			xa := float64(iCoords[0])
			ya := float64(iCoords[1])
			for j := i + 1; j < len(letterEmitters); j++ {
				jCoords := letterEmitters[j]
				xb := float64(jCoords[0])
				yb := float64(jCoords[1])

				distX := math.Abs(xb - xa)
				distY := math.Abs(yb - ya)

				for k := float64(1); ; k++ {
					distXK := distX * k
					distYK := distY * k
					xcFloat := xa - distXK
					if xa > xb {
						xcFloat = xa + distXK
					}
					ycFloat := ya - distYK
					if ya > yb {
						ycFloat = ya + distYK
					}

					xc := int(xcFloat)
					yc := int(ycFloat)

					if xc < len(lines) && xc >= 0 && yc < len(lines) && yc >= 0 {
						antinodes[[2]int{xc, yc}] = struct{}{}
					} else {
						break
					}
				}

				for k := float64(1); ; k++ {
					distXK := distX * k
					distYK := distY * k
					xdFloat := xb + distXK
					if xb < xa {
						xdFloat = xb - distXK
					}
					ydFloat := yb + distYK
					if yb < ya {
						ydFloat = yb - distYK
					}

					xd := int(xdFloat)
					yd := int(ydFloat)

					if xd < len(lines) && xd >= 0 && yd < len(lines) && yd >= 0 {
						antinodes[[2]int{xd, yd}] = struct{}{}
					} else {
						break
					}
				}
			}
		}
	}

	fmt.Println(len(antinodes))

	return nil
}
