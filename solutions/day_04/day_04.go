package day_04

import (
	"slices"
	"strings"
)

type direction struct {
	x int
	y int
}

var directions = map[string]direction{
	"n":  {x: -1, y: 0},
	"ne": {x: -1, y: 1},
	"e":  {x: 0, y: 1},
	"se": {x: 1, y: 1},
	"s":  {x: 1, y: 0},
	"sw": {x: 1, y: -1},
	"w":  {x: 0, y: -1},
	"nw": {x: -1, y: -1},
}

func setup(input string) [][]string {
	cleanInput, _ := strings.CutSuffix(input, "\n")
	lines := strings.Split(cleanInput, "\n")
	grid := make([][]string, len(lines))

	for i, l := range lines {
		grid[i] = strings.Split(l, "")
	}
	return grid
}

func checkAccessibility(p_x, p_y int, grid [][]string) bool {
	counter := 0
	max_x := len(grid[0])
	max_y := len(grid)

	for _, d := range directions {
		new_x := p_x - d.x
		new_y := p_y - d.y
		if new_x >= 0 && new_x < max_x && new_y >= 0 && new_y < max_y {
			if grid[new_y][new_x] == "@" {
				counter++
			}
		}
	}
	if counter >= 4 {
		return false
	}
	return true
}

func Part1(input string) (int, error) {
	result := 0
	grid := setup(input)

	for y := range len(grid[0]) {
		for x := range len(grid) {
			if grid[y][x] != "." {
				accessible := checkAccessibility(x, y, grid)
				if accessible {
					result++
				}
			}
		}
	}

	return result, nil
}

func getCounter(p_x, p_y int, grid [][]string) int {
	counter := 0
	max_x := len(grid[0])
	max_y := len(grid)

	for _, d := range directions {
		new_x := p_x - d.x
		new_y := p_y - d.y
		if new_x >= 0 && new_x < max_x && new_y >= 0 && new_y < max_y {
			if grid[new_y][new_x] == "@" {
				counter++
			}
		}
	}
	return counter
}

type point struct {
	x int
	y int
}

func removeRoll(p point, grid [][]int) []point {
	max_x := len(grid[0])
	max_y := len(grid)
	new_values := []point{}

	for _, d := range directions {
		new_x := p.x - d.x
		new_y := p.y - d.y
		if new_x >= 0 && new_x < max_x && new_y >= 0 && new_y < max_y {
			grid[new_y][new_x]--
			if grid[new_y][new_x] > 0 {
				new_values = append(new_values, point{new_x, new_y})
			}
		}
	}
	return new_values
}

func Part2(input string) (int, error) {
	result := 0
	grid := setup(input)
	counterGrid := make([][]int, len(grid))
	for i := range counterGrid {
		counterGrid[i] = make([]int, len(grid[i]))
	}

	queue := []point{}

	for y := range len(grid[0]) {
		for x := range len(grid) {
			if grid[y][x] != "." {
				counter := getCounter(x, y, grid)
				counterGrid[y][x] = counter
			} else {
				counterGrid[y][x] = -1
			}
			queue = append(queue, point{x, y})
		}
	}

	for len(queue) > 0 {
		crntPoint := queue[0]
		if counterGrid[crntPoint.y][crntPoint.x] < 4 && counterGrid[crntPoint.y][crntPoint.x] >= 0 {
			changedPoints := removeRoll(crntPoint, counterGrid)
			counterGrid[crntPoint.y][crntPoint.x] = -1
			queue = append(queue, changedPoints...)
			result++
		}
		queue = slices.Delete(queue, 0, 1)
	}

	return result, nil
}
