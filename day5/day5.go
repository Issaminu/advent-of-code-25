package main

import (
	"advent25"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := advent25.ReadInputOfDay(5)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	lines := strings.Split(string(input), "\n\n")

	ids := strings.Split(string(lines[0]), "\n")
	ingredients := strings.Split(string(lines[1]), "\n")

	var answer int

	var ranges = make([]advent25.Pair[int, int], len(ids))

	for i, line := range ids {
		idRange := strings.Split(string(line), "-")
		start, _ := strconv.Atoi(idRange[0])
		end, _ := strconv.Atoi(idRange[1])

		ranges[i] = advent25.Pair[int, int]{First: start, Second: end}

	}

	for _, ing := range ingredients {
		ing_int, _ := strconv.Atoi(ing)

		for _, idRange := range ranges {
			if idRange.First <= ing_int && idRange.Second >= ing_int {
				answer++
				break
			}
		}

	}

	println(answer)

}

func partTwo(input string) {
	lines := strings.Split(string(input), "\n\n")

	ids := strings.Split(string(lines[0]), "\n")

	var answer int

	// sorted array of elements, each element is a pair of (min,max)
	var ranges []advent25.Pair[int, int]

	// var visited = make(map[int]struct{})

	for _, line := range ids {
		idRange := strings.Split(string(line), "-")
		start, _ := strconv.Atoi(idRange[0])
		end, _ := strconv.Atoi(idRange[1])

		ranges = append(ranges, advent25.Pair[int, int]{First: start, Second: end})
	}

	slices.SortFunc(ranges, func(a, b advent25.Pair[int, int]) int {
		return a.First - b.First
	})

	for i := 1; i < len(ranges); i++ {
		prevRange := &ranges[i-1]
		if prevRange.Second > ranges[i].Second {

			// pop current range since the prev range already contains it
			ranges = slices.Delete(ranges, i, i+1)
			i--
			continue
		}

		if prevRange.Second >= ranges[i].First || ranges[i].First-prevRange.Second == 1 {
			prevRange.Second = ranges[i].Second
			// pop current range since the prev range got larger
			ranges = slices.Delete(ranges, i, i+1)
			i--
		}
	}

	for _, idRange := range ranges {
		answer += idRange.Second - idRange.First + 1
	}

	println(answer)

}
