package day_03

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func findJoltage(batteries []int) int {
	originalBatteries := make([]int, len(batteries))
	copy(originalBatteries, batteries)

	m1 := slices.Max(batteries[:len(batteries)-1])
	m1_i := slices.Index(originalBatteries, m1)
	batteries = slices.Delete(batteries, 0, m1_i+1)

	m2 := slices.Max(batteries)
	return m1*10 + m2
}

func sliceToInt(s []int) int {
	res := 0
	for i, x := range s {
		res += x * int(math.Pow(float64(10), float64(len(s)-i-1)))
	}
	return res
}

func findJoltage2(batteries []int) int {
	originalBatteries := make([]int, len(batteries))
	copy(originalBatteries, batteries)

	maxBatteries := 12
	batteriesToTurnOn := maxBatteries

	turnedOn := make([]int, batteriesToTurnOn)

	for batteriesToTurnOn > 0 {
		max_i := len(batteries) - batteriesToTurnOn + 1
		m1 := slices.Max(batteries[:max_i])
		m1_i := slices.Index(batteries, m1)
		batteries = slices.Delete(batteries, 0, m1_i+1)

		turnedOn[maxBatteries-batteriesToTurnOn] = m1
		batteriesToTurnOn--
	}

	res := sliceToInt(turnedOn)
	return res
}

func Part1(input string) (int, error) {
	result := 0

	banks := strings.Split(input, "\n")
	banks = banks[:len(banks)-1]
	for _, bank := range banks {
		batteriesInt := make([]int, len(banks[0]))
		batteries := strings.Split(bank, "")
		for i, b := range batteries {
			crntNum, err := strconv.Atoi(b)
			if err != nil {
				return 0, err
			}
			batteriesInt[i] = crntNum
		}
		x := findJoltage(batteriesInt)
		result += x
	}

	return result, nil
}

func Part2(input string) (int, error) {
	result := 0

	banks := strings.Split(input, "\n")
	banks = banks[:len(banks)-1]
	for _, bank := range banks {
		batteriesInt := make([]int, len(banks[0]))
		batteries := strings.Split(bank, "")
		for i, b := range batteries {
			crntNum, err := strconv.Atoi(b)
			if err != nil {
				return 0, err
			}
			batteriesInt[i] = crntNum
		}
		x := findJoltage2(batteriesInt)
		result += x
	}

	return result, nil
}
