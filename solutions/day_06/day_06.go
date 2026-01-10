package day_06

import (
	"strconv"
	"strings"
)

func setup(input string) ([][]int, []string) {
	cleanInput, _ := strings.CutSuffix(input, "\n")
	lines := strings.Split(cleanInput, "\n")
	numbersLines := lines[:len(lines)-1]
	mathOps := strings.Fields(lines[len(lines)-1])

	tempNumbersLines := make([][]int, len(numbersLines))

	for i, line := range numbersLines {
		nums := strings.Fields(line)
		temp := make([]int, len(mathOps))
		for j, num := range nums {
			num_, err := strconv.Atoi(num)
			if err != nil {
				return nil, nil
			}
			temp[j] = num_
		}
		tempNumbersLines[i] = temp
	}
	return tempNumbersLines, mathOps
}

func setup2(input string) ([][]int, []string) {
	cleanInput, _ := strings.CutSuffix(input, "\n")
	lines := strings.Split(cleanInput, "\n")
	numbersLines := lines[:len(lines)-1]
	mathOps := strings.Fields(lines[len(lines)-1])

	numbersCol := make([]string, len(numbersLines[0]))

	for i, _ := range numbersLines[0] {
		crntCol := ""
		for _, line := range numbersLines {
			crntCol += string(line[i])
		}
		numbersCol[i] = crntCol
	}

	numbers := [][]int{}
	crntNumbers := []int{}

	for _, nCol := range numbersCol {
		tempNums := strings.Fields(nCol)
		crntNumString := strings.Join(tempNums, "")
		if crntNumString != "" {
			num, err := strconv.Atoi(crntNumString)
			if err != nil {
				return nil, nil
			}
			crntNumbers = append(crntNumbers, num)
		} else {
			numbers = append(numbers, crntNumbers)
			crntNumbers = nil
		}
	}
	numbers = append(numbers, crntNumbers)
	return numbers, mathOps
}

func doMath(x, y int, op string) int {
	result := 0
	switch op {
	case "+":
		result = x + y
	case "*":
		result = x * y
	}
	return result
}

func Part1(input string) (int, error) {
	result := 0
	numberLines, mathOps := setup(input)
	ops := len(mathOps)
	results := numberLines[0]

	for line := 1; line < len(numberLines); line++ {
		for i := range ops {
			results[i] = doMath(results[i], numberLines[line][i], mathOps[i])
		}
	}

	for _, r := range results {
		result += r
	}
	return result, nil
}

func Part2(input string) (int, error) {
	result := 0
	numberCols, mathOps := setup2(input)

	for j, col := range numberCols {
		colResult := col[0]
		for i := 1; i < len(col); i++ {
			colResult = doMath(colResult, col[i], mathOps[j])
		}
		result += colResult
	}

	return result, nil
}
