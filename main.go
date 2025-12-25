package main

import (
	"fmt"

	"github.com/chickiexd/advent_of_code_2025/solutions/day_01"
	"github.com/chickiexd/advent_of_code_2025/solutions/day_02"
	"github.com/chickiexd/advent_of_code_2025/solutions/day_03"
	"github.com/chickiexd/advent_of_code_2025/solutions/day_04"
	"github.com/chickiexd/advent_of_code_2025/utils"
)

const day = 4

// const part1_test_input = true
const part1_test_input = false

// const part2_test_input = true
const part2_test_input = false

var funcs = map[int][]func(string) (int, error){
	1: {day_01.Part1, day_01.Part2},
	2: {day_02.Part1, day_02.Part2},
	3: {day_03.Part1, day_03.Part2},
	4: {day_04.Part1, day_04.Part2},
}

func main() {
	testInput := utils.GetTestInput(day)
	realInput := utils.GetInput(day)

	var input string
	if part1_test_input {
		input = testInput
	} else {
		input = realInput
	}
	res := runPart(1, input)
	fmt.Println("Part 1 Result:")
	fmt.Println(res)

	if part2_test_input {
		input = testInput
	} else {
		input = realInput
	}
	res = runPart(2, input)
	fmt.Println("Part 2 Result:")
	fmt.Println(res)
}

func runPart(part int, input string) int {
	res, err := funcs[day][part-1](input)
	if err != nil {
		panic(err)
	}
	return res
}
