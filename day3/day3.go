package main

import (
	"advent25"
	"strconv"
	"strings"
)

func main() {
	input := advent25.ReadInputOfDay(3)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	lines := strings.Split(string(input), "\n")
	var answer int
	for _, line := range lines {
		batteries := make([]int, len(line))

		var biggest int
		var biggestIdx int
		for i := range line {
			battery, _ := strconv.Atoi(string(line[i]))
			batteries[i] = battery
			if i != len(line)-1 {
				if battery > biggest {
					biggest = battery
					biggestIdx = i
				}
			}
		}

		var secondBiggest int
		for i := biggestIdx + 1; i < len(batteries); i++ {
			secondBiggest = max(secondBiggest, batteries[i])
		}

		answer += biggest*10 + secondBiggest
	}
	println(answer)
}

func partTwo(input string) {
	lines := strings.Split(string(input), "\n")
	var answer int
	for _, line := range lines {
		batteries := make([]int, len(line))

		for i := range line {
			battery, _ := strconv.Atoi(string(line[i]))
			batteries[i] = battery
		}

		var results [12]int

		start := 0
		for bat := range 12 {
			biggest := batteries[start]
			biggestIdx := start
			for i := start; i <= len(batteries)-1; i++ {
				if len(batteries)-i < 12-bat {
					break
				}

				if batteries[i] > biggest {
					biggest = batteries[i]
					biggestIdx = i
				}

			}
			start = biggestIdx + 1
			results[bat] = biggest
		}

		var max int
		for _, result := range results {
			max = max*10 + result
		}

		answer += max
	}
	println(answer)
}
