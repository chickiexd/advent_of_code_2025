package day_02

import (
	"strconv"
	"strings"
)

func checkInvalid(num int) bool {
	numStr := strconv.Itoa(num)
	if len(numStr)%2 != 0 {
		return false
	}

	l := 0
	r := len(numStr) / 2
	for l < (len(numStr) / 2) {
		if numStr[l] != numStr[r] {
			return false
		}
		l++
		r++
	}
	return true
}

func checkSeq(s string, l int) bool {
	seqAmount := len(s) / l
	crntSeq := s[:l]
	for i := 1; i < seqAmount; i++ {
		if crntSeq != s[i*l:(i+1)*l] {
			return false
		}
	}
	return true
}

func checkInvalid2(num int) bool {
	numStr := strconv.Itoa(num)

	possibleSeqLengths := []int{}

	for i := 1; i < len(numStr); i++ {
		if len(numStr)%i == 0 {
			possibleSeqLengths = append(possibleSeqLengths, i)
		}
	}

	for _, seqLen := range possibleSeqLengths {
		valid := checkSeq(numStr, seqLen)
		if valid {
			return true
		}
	}
	return false
}

func Part1(input string) (int, error) {
	result := 0

	cleanInput, _ := strings.CutSuffix(input, "\n")
	for vr := range strings.SplitSeq(cleanInput, ",") {
		fromTo := strings.Split(vr, "-")
		from, err := strconv.Atoi(fromTo[0])
		to, err := strconv.Atoi(fromTo[1])
		if err != nil {
			return 0, err
		}

		for i := from; i <= to; i++ {
			valid := checkInvalid(i)
			if valid {
				result += i
			}
		}
	}

	return result, nil
}

func Part2(input string) (int, error) {
	result := 0

	cleanInput, _ := strings.CutSuffix(input, "\n")
	for vr := range strings.SplitSeq(cleanInput, ",") {
		fromTo := strings.Split(vr, "-")
		from, err := strconv.Atoi(fromTo[0])
		to, err := strconv.Atoi(fromTo[1])
		if err != nil {
			return 0, err
		}

		for i := from; i <= to; i++ {
			valid := checkInvalid2(i)
			if valid {
				result += i
			}
		}
	}

	return result, nil
}
