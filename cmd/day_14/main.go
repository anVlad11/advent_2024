package main

import (
	"fmt"
	"github.com/anVlad11/advent_2024/pkg/utils"
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func do() error {
	input, err := utils.GetInput("inputs/day_14/input_1.txt")
	if err != nil {
		return err
	}

	//err = part1(input, 11, 7, 100, 100)
	err = part1(input, 101, 103, 100)
	if err != nil {
		return err
	}

	err = part2FloodFill(input, 101, 103)
	//err = part2(input, 11, 7)
	//err = part2(input, 101, 103)
	if err != nil {
		return err
	}

	return nil
}

var regex = regexp.MustCompile(`p=(.*),(.*) v=(.*),(.*)$`)

func part1(lines []string, maxX int, maxY int, steps int) error {

	matrix := make([][]int, maxY)
	for j := range matrix {
		matrix[j] = make([]int, maxX)
	}

	safetyQuadrants := map[[2]int]int{
		[2]int{-1, -1}: 0,
		[2]int{-1, 1}:  0,
		[2]int{1, -1}:  0,
		[2]int{1, 1}:   0,
	}

	for _, line := range lines {
		res := regex.FindAllStringSubmatch(line, -1)
		p := [2]int{utils.MustParseInt(res[0][1]), utils.MustParseInt(res[0][2])}
		v := [2]int{utils.MustParseInt(res[0][3]), utils.MustParseInt(res[0][4])}

		final := [2]int{(p[0] + (v[0] * steps)) % maxX, (p[1] + (v[1] * steps)) % maxY}
		if final[0] < 0 {
			final[0] += maxX
		}
		if final[1] < 0 {
			final[1] += maxY
		}

		safetyQuadrant := [2]int{0, 0}
		if final[0] > maxX/2 {
			safetyQuadrant[0] = 1
		}
		if final[0] < maxX/2 {
			safetyQuadrant[0] = -1
		}
		if final[1] > maxY/2 {
			safetyQuadrant[1] = 1
		}
		if final[1] < maxY/2 {
			safetyQuadrant[1] = -1
		}
		if safetyQuadrant[0] != 0 && safetyQuadrant[1] != 0 {
			safetyQuadrants[safetyQuadrant]++
		}

		matrix[final[1]][final[0]]++
	}

	sum := 1
	for _, v := range safetyQuadrants {
		sum *= v
	}
	fmt.Println(sum)

	return nil
}

// part2 requires you to visually inspect the outputs.
func part2(lines []string, maxX int, maxY int) error {
	stepsStart := 0
	stepsStop := maxX * maxY

	for i := stepsStart; i <= stepsStop; i++ {
		img := image.NewRGBA(image.Rect(0, 0, maxX, maxY))
		for x := range img.Rect.Dx() {
			for y := range img.Rect.Dy() {
				img.Set(x, y, color.Black)
			}
		}
		matrix := make([][]int, maxY)
		for j := range matrix {
			matrix[j] = make([]int, maxX)
		}

		for _, line := range lines {
			res := regex.FindAllStringSubmatch(line, -1)
			p := [2]int{utils.MustParseInt(res[0][1]), utils.MustParseInt(res[0][2])}
			v := [2]int{utils.MustParseInt(res[0][3]), utils.MustParseInt(res[0][4])}

			final := [2]int{(p[0] + (v[0] * i)) % maxX, (p[1] + (v[1] * i)) % maxY}
			if final[0] < 0 {
				final[0] += maxX
			}
			if final[1] < 0 {
				final[1] += maxY
			}

			matrix[final[1]][final[0]]++
			img.Set(final[0], final[1], color.White)
		}
		/*
			fmt.Printf("After %d second:\n", i)
			res := []string{}
			for y, _ := range matrix {
				for _, sym := range matrix[y] {
					if sym == 0 {
						res = append(res, ".")
					} else {
						res = append(res, fmt.Sprintf("%d", sym))
					}
				}
				res = append(res, "\n")
			}
			fmt.Printf("%s\n", strings.Join(res, ""))

		*/
		filename := fmt.Sprintf("pkg/data/outputs/day_14/step_%d.png", i)
		file, err := os.Create(filename)
		if err != nil {
			return err
		}

		err = png.Encode(file, img)
		if err != nil {
			return err
		}

		file.Close()
	}

	return nil
}

var directions = [][2]int{{0, -1}, {1, 0}, {-1, 0}, {0, 1}}

// part2FloodFill flood fill approach is derived
// from knowing the solution in advance.
// I didn't like this puzzle.
func part2FloodFill(lines []string, maxX int, maxY int) error {
	stepsStart := 0
	stepsStop := maxX * maxY

	maxSize := 0
	maxSizeStep := 0
	for i := stepsStart; i <= stepsStop; i++ {
		matrix := make([][]int, maxY)
		for j := range matrix {
			matrix[j] = make([]int, maxX)
		}

		for _, line := range lines {
			res := regex.FindAllStringSubmatch(line, -1)
			p := [2]int{utils.MustParseInt(res[0][1]), utils.MustParseInt(res[0][2])}
			v := [2]int{utils.MustParseInt(res[0][3]), utils.MustParseInt(res[0][4])}

			final := [2]int{(p[0] + (v[0] * i)) % maxX, (p[1] + (v[1] * i)) % maxY}
			if final[0] < 0 {
				final[0] += maxX
			}
			if final[1] < 0 {
				final[1] += maxY
			}

			matrix[final[1]][final[0]]++
		}

		visited := map[[2]int]struct{}{}
		for x := 0; x < maxX; x++ {
			for y := 0; y < maxY; y++ {
				if _, exists := visited[[2]int{x, y}]; exists {
					continue
				}

				size := 0
				queue := [][2]int{{x, y}}
				for len(queue) > 0 {
					item := queue[0]
					queue = queue[1:]

					if item[0] < 0 || item[0] >= maxX || item[1] < 0 || item[1] >= maxY {
						continue
					}

					if _, exists := visited[[2]int{item[0], item[1]}]; exists {
						continue
					}
					visited[[2]int{item[0], item[1]}] = struct{}{}

					if matrix[item[1]][item[0]] <= 0 {
						continue
					}

					size++
					for _, direction := range directions {
						queue = append(queue, [2]int{item[0] + direction[0], item[1] + direction[1]})
					}
				}

				if size > 0 {
					if size > maxSize {
						maxSize = size
						maxSizeStep = i
					}
				}
			}
		}
	}

	fmt.Println(maxSizeStep)

	return nil
}
