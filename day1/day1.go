package main

import (
	"advent25"
	"strconv"
	"strings"
)

func main() {
	input := advent25.ReadInputOfDay(1)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	lines := strings.Split(string(input), "\n")
	answer := 0
	position := 50
	for _, line := range lines {
		move := line[0]
		distance, _ := strconv.Atoi(line[1:])

		if move == 'L' {
			position -= distance
		} else {
			position += distance
		}
		if position%100 == 0 {
			answer++
		}
	}

	println(answer)
}

func partTwo(input string) {
	lines := strings.Split(string(input), "\n")
	answer := 0
	position := 50

	for _, line := range lines {
		move := line[0]
		distance, _ := strconv.Atoi(line[1:])

		if move == 'R' {
			// incredible engineering happening here
			for curr := position; curr < position+distance; curr++ {
				if curr%100 == 0 {
					answer++
				}
			}
			position += distance
		} else {
			for curr := position; curr > position-distance; curr-- {
				if curr%100 == 0 {
					answer++
				}
			}
			position -= distance
		}
	}

	println(answer)
}
