package day_01

import (
	"strconv"
	"strings"
)

func Part1(input string) (int, error) {
	commands := strings.Split(input, "\n")
	commands = commands[:len(commands)-1]

	dial := 50
	res := 0
	for _, c := range commands {
		amount := c[1:]
		a, err := strconv.Atoi(amount)
		if err != nil {
			return 0, err
		}
		if c[0] == 'L' {
			dial -= a
		}
		if c[0] == 'R' {
			dial += a
		}
		dial = (dial%100 + 100) % 100

		if dial == 0 {
			res++
		}
	}
	return res, nil
}

func Part2(input string) (int, error) {
	commands := strings.Split(input, "\n")
	commands = commands[:len(commands)-1]

	dial := 50
	res := 0
	for _, c := range commands {
		amount := c[1:]
		a, err := strconv.Atoi(amount)
		if err != nil {
			return 0, err
		}

		res += a / 100
		a = a % 100

		if c[0] == 'L' {
			if a > dial && dial != 0 {
				res++
			}
			dial -= a
		}
		if c[0] == 'R' {
			if dial+a > 100 {
				res++
			}
			dial += a
		}
		dial = (dial%100 + 100) % 100

		if dial == 0 {
			res++
		}
	}
	return res, nil
}
