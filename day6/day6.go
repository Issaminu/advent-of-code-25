package main

import (
	"advent25"
	"strconv"
	"strings"
)

func main() {
	input := advent25.ReadInputOfDay(6)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	rawLines := strings.Split(input, "\n")
	lines := make([][]string, len(rawLines))

	for i := range lines {
		cleanedLine := cleanLine(rawLines[i])
		lines[i] = strings.Split(cleanedLine, " ")
	}
	var answer int

	for i := 0; i < len(lines[0]); i++ {
		localAnswer, _ := strconv.Atoi(lines[0][i])
		operation := lines[len(lines)-1][i]

		for j := 1; j < len(lines)-1; j++ {

			num, _ := strconv.Atoi(lines[j][i])
			switch operation {
			case "+":
				localAnswer += num
			case "*":
				localAnswer *= num
			}
		}
		answer += localAnswer
	}

	println(answer)
}

func partTwo(input string) {
	lines := strings.Split(input, "\n")
	var answer int

	var operation byte

	var nums []int

	for i := len(lines[0]) - 1; i >= 0; i-- {
		var temp string
		for j := 0; j < len(lines)-1; j++ {
			if lines[j][i] != ' ' {
				temp += string(lines[j][i])
			}
		}
		val, _ := strconv.Atoi(temp)
		if val != 0 {
			nums = append(nums, val)
		}

		if lines[len(lines)-1][i] != ' ' {
			operation = lines[len(lines)-1][i]
			localAns := 0
			if operation == '*' {
				localAns = 1
			}
			for _, num := range nums {
				switch operation {
				case '+':
					localAns += num
				case '*':
					localAns *= num
				default:
					panic("hehe")
				}
			}
			answer += localAns
			nums = []int{}
			operation = ' '
		}
	}

	println(answer)
}

func cleanLine(line string) string {
	line = strings.TrimSpace(line)
	var newStr string
	for i := 0; i < len(line); i++ {
		if line[i] != ' ' {
			newStr += string(line[i])
		} else if newStr[len(newStr)-1] != ' ' {
			newStr += string(' ')
		}
	}
	return newStr
}
