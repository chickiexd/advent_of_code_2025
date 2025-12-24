package day_02

import (
	"testing"

	"github.com/chickiexd/advent_of_code_2025/utils"
)

func TestPart1(t *testing.T) {
	my_result, err := Part1(utils.GetInput(1))
	if err != nil {
		t.Errorf("Part1() returned an error: %v", err)
	}
	expected_result := 1

	if my_result != expected_result {
		t.Errorf("Part1() = %d; want %d", my_result, expected_result)
	}
}

func TestPart2(t *testing.T) {
	my_result, err := Part2(utils.GetInput(1))
	if err != nil {
		t.Errorf("Part1() returned an error: %v", err)
	}
	expected_result := 1

	if my_result != expected_result {
		t.Errorf("Part1() = %d; want %d", my_result, expected_result)
	}
}

// func TestGetInput(t *testing.T) {
// 	input := utils.GetInput(1)
// 	expected_input := "PLACEHOLDER FOR DAY INPUT"
//
// 	if input != expected_input {
// 		t.Errorf("GetInput(1) = %s; want %s", input, expected_input)
// 	}
// }
