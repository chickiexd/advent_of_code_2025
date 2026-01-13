package day_07

import (
	"errors"
	"slices"
	"strings"

	"github.com/chickiexd/advent_of_code_2025/utils"
)

func setup(input string) [][]string {
	cleanInput, _ := strings.CutSuffix(input, "\n")
	lines := strings.Split(cleanInput, "\n")

	grid := make([][]string, len(lines))
	for i, l := range lines {
		lineSplit := strings.Split(l, "")
		grid[i] = lineSplit
	}

	return grid
}

type Point struct {
	x int
	y int
}

func getStart(grid [][]string) (Point, error) {
	for y := range len(grid) {
		for x := range len(grid[y]) {
			if grid[y][x] == "S" {
				return Point{x, y}, nil
			}
		}
	}
	return Point{}, errors.New("lul")
}

func inBounds(p Point, max_x, max_y int) bool {
	if p.x >= 0 && p.x < max_x && p.y >= 0 && p.y < max_y {
		return true
	}
	return false
}

func moveDown(p Point, grid [][]string) (bool, []Point) {
	downP := Point{p.x, p.y + 1}
	if inBounds(downP, len(grid[0]), len(grid)) == false {
		return false, nil
	}

	y := downP.y
	x := downP.x

	if grid[y][x] == "^" {
		newPointLeft := Point{x - 1, y}
		newPointRight := Point{x + 1, y}
		newPoints := []Point{newPointLeft, newPointRight}
		return true, newPoints
	} else {
		return false, []Point{downP}
	}
}

func moveDown2(p Point_, grid [][]string) (bool, []Point_) {
	downP_ := Point_{p.x, p.y + 1, p.d}
	if inBounds(Point{downP_.x, downP_.y}, len(grid[0]), len(grid)) == false {
		return false, nil
	}

	y := downP_.y
	x := downP_.x

	if grid[y][x] == "^" {
		newPointLeft := Point_{x - 1, y, p.d}
		newPointRight := Point_{x + 1, y, p.d}
		newPoints := []Point_{newPointLeft, newPointRight}
		return true, newPoints
	} else {
		return false, []Point_{downP_}
	}
}

func Part1(input string) (int, error) {
	result := 0
	grid := setup(input)
	max_y := len(grid)
	max_x := len(grid[0])

	startPoint, err := getStart(grid)
	if err != nil {
		return 0, err
	}

	q := utils.Q[Point]{}
	q.Put(startPoint)
	seen := []Point{}

	for q.Len() > 0 {
		crntPoint, err := q.Pop()
		seen = append(seen, crntPoint)
		if err != nil {
			return result, err
		}

		split, newPoints := moveDown(crntPoint, grid)
		if split {
			result++
		}
		for _, newP := range newPoints {
			if inBounds(newP, max_x, max_y) == true {
				if slices.Index(seen, newP) == -1 && slices.Index(q.Data, newP) == -1 {
					q.Put(newP)
				}
			}
		}
	}

	return result, nil
}

type Point_ struct {
	x int
	y int
	d int
}

func Part2(input string) (int, error) {
	result := 0
	grid := setup(input)
	max_y := len(grid)
	max_x := len(grid[0])

	startPoint, err := getStart(grid)
	if err != nil {
		return 0, err
	}

	q := utils.Q[Point_]{}

	q.Put(Point_{startPoint.x, startPoint.y, 1})
	seen := []Point_{}

	for q.Len() > 0 {
		crntPoint, err := q.Pop()
		seen = append(seen, crntPoint)
		if crntPoint.y == max_y-1 {
			result += crntPoint.d
		}
		if err != nil {
			return result, err
		}

		_, newPoints := moveDown2(crntPoint, grid)
		for _, newP := range newPoints {

			if inBounds(Point{newP.x, newP.y}, max_x, max_y) == true {
				if idx := pointInQ(newP, q); idx != -1{
					q.Data[idx].d += newP.d
				} else if slices.Index(seen, newP) == -1 {
					q.Put(newP)
				}
			}
		}
	}

	return result, nil
}

func pointInQ(p Point_, q utils.Q[Point_]) int {
	for i, x := range q.Data {
		if x.x == p.x && x.y == p.y {
			return i
		}
	}
	return -1
}
