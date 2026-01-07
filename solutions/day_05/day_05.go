package day_05

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/chickiexd/advent_of_code_2025/utils"
)

func setup(input string) ([]string, []string) {
	cleanInput, _ := strings.CutSuffix(input, "\n")
	parts := strings.Split(cleanInput, "\n\n")
	return strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")
}

func Part1(input string) (int, error) {
	result := 0
	ranges, ingredients := setup(input)
	for len(ingredients) > 0 {
		ingredient, err := strconv.Atoi(ingredients[0])
		if err != nil {
			return 0, err
		}
		for r := range ranges {
			min_max := strings.Split(ranges[r], "-")
			min_, err := strconv.Atoi(min_max[0])
			max_, err := strconv.Atoi(min_max[1])
			if err != nil {
				return 0, err
			}
			if ingredient >= min_ && ingredient <= max_ {
				result++
				break
			}
		}
		ingredients = slices.Delete(ingredients, 0, 1)
	}
	return result, nil
}

func makeRangesInt(input []string) [][]int {
	ranges_int := [][]int{}
	for _, r := range input {
		crnt := strings.Split(r, "-")
		min_, err := strconv.Atoi(crnt[0])
		max_, err := strconv.Atoi(crnt[1])
		if err != nil {
			return nil
		}
		new_r := []int{min_, max_}
		ranges_int = append(ranges_int, new_r)
	}
	return ranges_int
}

func Part2(input string) (int, error) {
	result := 0
	ranges, _ := setup(input)

	ranges_int := makeRangesInt(ranges)

	exclusive_ranges := [][]int{}

	sort.Slice(ranges_int, func(i, j int) bool {
		return ranges_int[i][0] < ranges_int[j][0]
	})

	q := utils.Q[[]int]{}
	q.Put(ranges_int[0])

	ranges_int = slices.Delete(ranges_int, 0, 1)

	for q.Len() > 0 {
		crnt_range, err := q.Pop()
		if err != nil {
			return result, err
		}

		crnt_min := crnt_range[0]
		crnt_max := crnt_range[1]

		new_ranges := [][]int{}
		for _, r := range ranges_int {
			r_min := r[0]
			r_max := r[1]

			if (r_min <= crnt_min && r_max >= crnt_min) || (r_max >= crnt_max && r_min <= crnt_max) || (r_min >= crnt_min && r_max <= crnt_max) {
				if r_min <= crnt_min && r_max >= crnt_min {
					crnt_min = r_min
				}
				if r_max >= crnt_max && r_min <= crnt_max {
					crnt_max = r_max
				}
			} else {
				new_ranges = append(new_ranges, r)
			}
		}
		new_range := []int{crnt_min, crnt_max}
		if len(ranges_int) == len(new_ranges) {
			exclusive_ranges = append(exclusive_ranges, new_range)
			if len(ranges_int) > 0 {
				q.Put(ranges_int[0])
				ranges_int = slices.Delete(ranges_int, 0, 1)
			}
		} else {
			q.Put(new_range)
			ranges_int = new_ranges
		}
	}
	for _, e := range exclusive_ranges {
		result += (e[1] - e[0]) + 1
	}

	return result, nil
}
