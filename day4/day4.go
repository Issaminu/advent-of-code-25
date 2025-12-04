package main

import (
	"advent25"
	"strings"
)

func main() {
	input := advent25.ReadInputOfDay(4)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	linesStr := strings.Split(string(input), "\n")
	lines := make([][]byte, len(linesStr))
	for i, s := range linesStr {
		lines[i] = []byte(s)
	}
	answer := 0
	for i := range lines {
	outer:
		for j := range lines[i] {
			if advent25.SafeGetElement(lines, i, j, 'A') != '@' {
				continue
			}

			directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}

			adjacents := 0
			for k := range directions {
				if advent25.SafeGetElement(lines, i+directions[k][0], j+directions[k][1], 'A') == '@' {
					adjacents++

					if adjacents >= 4 {
						continue outer
					}
				}
			}
			answer++

		}
	}
	println(answer)
}

func partTwo(input string) {
	linesStr := strings.Split(string(input), "\n")
	lines := make([][]byte, len(linesStr))
	for i, s := range linesStr {
		lines[i] = []byte(s)
	}
	answer := 0
	prevAnswer := -1
	for prevAnswer != answer {
		prevAnswer = answer
		for i := range lines {
			for j := range lines[i] {

				if advent25.SafeGetElement(lines, i, j, 'A') != '@' {
					continue
				}

				directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}

				adjacents := 0
				for k := range directions {
					if advent25.SafeGetElement(lines, i+directions[k][0], j+directions[k][1], 'A') == '@' {
						adjacents++
					}
				}

				if adjacents < 4 {
					lines[i][j] = '.'
					answer++
				}
			}
		}
	}
	println(answer)
}
