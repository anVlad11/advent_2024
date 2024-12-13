package main

import (
	"fmt"
	"github.com/anVlad11/advent_2024/pkg/utils"
	"regexp"
)

func main() {
	err := do()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func do() error {
	input, err := utils.GetInput("inputs/day_13/input_1.txt")
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

var regex = regexp.MustCompile(`.*X.(\d*), Y.(\d*)$`)

func part1(lines []string) error {
	// Just to simplify the type conversions later
	lines = append(lines, "\n")
	buttonA := [2]float64{}
	buttonB := [2]float64{}
	prize := [2]float64{}

	sum := int64(0)

	for i := 0; i < len(lines); i++ {
		localI := i % 4
		if localI == 3 {
			/*
				You need to press button A N times and button B M times to get to the prize at X Y.

				Let's look at the first block in the example:
				Button A: X+94, Y+34
				Button B: X+22, Y+67
				Prize: X=8400, Y=5400

				Each time you press button A, you increment X by 94
				Each time you press button B, you increment X by 22
				So, to get to the X = 8400 you need to increment N 94 times and M 22 times:
				94n + 22m = 8400

				Same idea for Y:
				34n + 67m = 5400

				Now you have a simple system of equations with two variables:
				{ 94n + 22m = 8400 }
				{ 34n + 67m = 5400 }

				Let's try to eliminate one of the variables here - the amount of button A presses - N.
				Multiply both sides of each equation by the amount of button A presses you need to do
				to equalize the multiplier for N in both equations.
				In this case, I will multiply the first equation by 34 and the second one by 94:

				{ 34*(94n + 22m) = 34 * 8400 } => { 3196x + 758y = 285600 }
				{ 94*(34n + 67m) = 94 * 5400 } => { 3196x + 6298y = 507600 }

				Nothing bad will happen if we subtract the left part of the second equation from the left part
				of the first equation and do the same for the right part:

				(3196n + 6298m) - (3196n + 758m) = 507600 - 285600
				0n + 5550m = 222000
				5550m = 222000
				m = 222000/5550
				m = 40

				Now we have the amount of the button B presses to reach the prize X - 40.
				As this amount is an integer, we don't need to throw out the whole block.

				Let's find out the amount of button A presses.
				For that, we can look at any of the original equations - i'll go with the first one -
				and substitute variable M in it with 40:
				94n + 22m = 8400
				94n + 22 * 40 = 8400
				94n = 8400 - 22*40
				94n = 8400 - 880
				94n = 7520
				n = 80

				This one is an integer too, so we can confirm that this block is solvable.

				Additional operations would include making sure that each button is pressed no more than 100 times
				and getting the amount of tokens for the solution - in this case, it's:
				80 * 3 + 40 * 1 = 280

				Now we need to represent this solution process in the imperative form.
				Let's go back to the original system of equations right after we added constant multipliers to it.

				{ 34*(94n + 22m) = 34 * 8400 } => { 3196x + 758y = 285600 }
				{ 94*(34n + 67m) = 94 * 5400 } => { 3196x + 6298y = 507600 }

				(3196n + 6298m) - (3196n + 758m) = 507600 - 285600
				3196n + 6298m - 3196n - 758m = 507600 - 285600

				5550y = 222000

				That becomes:
				{ buttonA[1] * (buttonA[0]*n + buttonB[0]*m) = buttonA[1] * prize[0] }
				{ buttonA[0] * (buttonA[1]*n + buttonB[1]*m) = buttonA[0] * prize[1] }

				(buttonA[0] * buttonB[1] * m) - (buttonA[1] * buttonB[0] * m) = (buttonA[0] * prize[1]) - (buttonA[1] * prize[0])

				((buttonA[0] * buttonB[1]) - (buttonA[1] * buttonB[0])) * m = (buttonA[0] * prize[1]) - (buttonA[1] * prize[0])

				m = ((buttonA[0] * prize[1]) - (buttonA[1] * prize[0])) / ((buttonA[0] * buttonB[1]) - (buttonA[1] * buttonB[0]))

				Finding N becomes simple too:
				94n + 22m = 8400
				buttonA[0] * n = buttonB[0] * m = prize[0]
				buttonA[0] * n = prize[0] - (buttonB[0] * m)
				n = (prize[0] - (buttonB[0] * m)) / buttonA[0]
			*/

			m := ((buttonA[0] * prize[1]) - (buttonA[1] * prize[0])) / ((buttonA[0] * buttonB[1]) - (buttonA[1] * buttonB[0]))
			if m > 100 || m != float64(int64(m)) {
				continue
			}

			n := (prize[0] - (buttonB[0] * m)) / buttonA[0]

			if n > 100 || n != float64(int64(n)) {
				continue
			}

			sum += int64(n*3 + m)
			//fmt.Printf("%v, %v, %v = %v %v\n", buttonA, buttonB, prize, n, m)

			continue
		}
		match := regex.FindAllStringSubmatch(lines[i], -1)
		x, y := utils.MustParseFloat64(match[0][1]), utils.MustParseFloat64(match[0][2])
		switch localI {
		case 0:
			buttonA = [2]float64{x, y}
		case 1:
			buttonB = [2]float64{x, y}
		case 2:
			prize = [2]float64{x, y}
		}
	}

	fmt.Println(sum)
	return nil
}

// part2 differences are minimal and do not increase complexity over part1
func part2(lines []string) error {
	// Just to simplify the type conversions later
	lines = append(lines, "\n")
	buttonA := [2]float64{}
	buttonB := [2]float64{}
	prize := [2]float64{}

	sum := int64(0)

	for i := 0; i < len(lines); i++ {
		localI := i % 4
		if localI == 3 {
			/*
				You need to press button A N times and button B M times to get to the prize at X Y.

				Let's look at the first block in the example:
				Button A: X+94, Y+34
				Button B: X+22, Y+67
				Prize: X=8400, Y=5400

				Each time you press button A, you increment X by 94
				Each time you press button B, you increment X by 22
				So, to get to the X = 8400 you need to increment N 94 times and M 22 times:
				94n + 22m = 8400

				Same idea for Y:
				34n + 67m = 5400

				Now you have a simple system of equations with two variables:
				{ 94n + 22m = 8400 }
				{ 34n + 67m = 5400 }

				Let's try to eliminate one of the variables here - the amount of button A presses - N.
				Multiply both sides of each equation by the amount of button A presses you need to do
				to equalize the multiplier for N in both equations.
				In this case, I will multiply the first equation by 34 and the second one by 94:

				{ 34*(94n + 22m) = 34 * 8400 } => { 3196x + 758y = 285600 }
				{ 94*(34n + 67m) = 94 * 5400 } => { 3196x + 6298y = 507600 }

				Nothing bad will happen if we subtract the left part of the second equation from the left part
				of the first equation and do the same for the right part:

				(3196n + 6298m) - (3196n + 758m) = 507600 - 285600
				0n + 5550m = 222000
				5550m = 222000
				m = 222000/5550
				m = 40

				Now we have the amount of the button B presses to reach the prize X - 40.
				As this amount is an integer, we don't need to throw out the whole block.

				Let's find out the amount of button A presses.
				For that, we can look at any of the original equations - i'll go with the first one -
				and substitute variable M in it with 40:
				94n + 22m = 8400
				94n + 22 * 40 = 8400
				94n = 8400 - 22*40
				94n = 8400 - 880
				94n = 7520
				n = 80

				This one is an integer too, so we can confirm that this block is solvable.

				Additional operations would include making sure that each button is pressed no more than 100 times
				and getting the amount of tokens for the solution - in this case, it's:
				80 * 3 + 40 * 1 = 280

				Now we need to represent this solution process in the imperative form.
				Let's go back to the original system of equations right after we added constant multipliers to it.

				{ 34*(94n + 22m) = 34 * 8400 } => { 3196x + 758y = 285600 }
				{ 94*(34n + 67m) = 94 * 5400 } => { 3196x + 6298y = 507600 }

				(3196n + 6298m) - (3196n + 758m) = 507600 - 285600
				3196n + 6298m - 3196n - 758m = 507600 - 285600

				5550y = 222000

				That becomes:
				{ buttonA[1] * (buttonA[0]*n + buttonB[0]*m) = buttonA[1] * prize[0] }
				{ buttonA[0] * (buttonA[1]*n + buttonB[1]*m) = buttonA[0] * prize[1] }

				(buttonA[0] * buttonB[1] * m) - (buttonA[1] * buttonB[0] * m) = (buttonA[0] * prize[1]) - (buttonA[1] * prize[0])

				((buttonA[0] * buttonB[1]) - (buttonA[1] * buttonB[0])) * m = (buttonA[0] * prize[1]) - (buttonA[1] * prize[0])

				m = ((buttonA[0] * prize[1]) - (buttonA[1] * prize[0])) / ((buttonA[0] * buttonB[1]) - (buttonA[1] * buttonB[0]))

				Finding N becomes simple too:
				94n + 22m = 8400
				buttonA[0] * n = buttonB[0] * m = prize[0]
				buttonA[0] * n = prize[0] - (buttonB[0] * m)
				n = (prize[0] - (buttonB[0] * m)) / buttonA[0]
			*/

			m := ((buttonA[0] * prize[1]) - (buttonA[1] * prize[0])) / ((buttonA[0] * buttonB[1]) - (buttonA[1] * buttonB[0]))
			if m != float64(int64(m)) {
				continue
			}

			n := (prize[0] - (buttonB[0] * m)) / buttonA[0]

			if n != float64(int64(n)) {
				continue
			}

			sum += int64(n*3 + m)
			//fmt.Printf("%v, %v, %v = %v %v\n", buttonA, buttonB, prize, n, m)

			continue
		}
		match := regex.FindAllStringSubmatch(lines[i], -1)
		x, y := utils.MustParseFloat64(match[0][1]), utils.MustParseFloat64(match[0][2])
		switch localI {
		case 0:
			buttonA = [2]float64{x, y}
		case 1:
			buttonB = [2]float64{x, y}
		case 2:
			prize = [2]float64{x + 10000000000000, y + 10000000000000}
		}
	}

	fmt.Println(sum)
	return nil
}
