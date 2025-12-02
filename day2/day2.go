package main

import (
	"advent25"
	"strconv"
	"strings"
)

func main() {
	input := advent25.ReadInputOfDay(2)
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	lines := strings.Split(string(input), ",")
	answer := 0

	for _, line := range lines {
		idRange_str := strings.Split(line, "-")
		start_str := idRange_str[0]
		end_str := idRange_str[1]

		start, _ := strconv.Atoi(start_str)
		end, _ := strconv.Atoi(end_str)

		for num := start; num <= end; num++ {
			num_str := strconv.Itoa(num)

			len := len(num_str)

			if len%2 == 1 {
				continue
			}

			if num_str[0:len/2] == num_str[(len/2):len] {
				answer += num
			}
		}
	}
	println(answer)
}

func partTwo(input string) {
	lines := strings.Split(string(input), ",")
	answer := 0

	for _, line := range lines {
		idRange_str := strings.Split(line, "-")
		start_str := idRange_str[0]
		end_str := idRange_str[1]

		start, _ := strconv.Atoi(start_str)
		end, _ := strconv.Atoi(end_str)

		for num := start; num <= end; num++ {
			num_str := strconv.Itoa(num)

			lenn := len(num_str)

			for i := 1; i <= lenn/2; i++ {
				if lenn%i == 0 {
					bad := false
					for j := 0; j < lenn; j += i {
						if num_str[0:i] != num_str[j:j+i] {
							bad = true
							break
						}
					}
					if !bad {
						answer += num
						break
					}
				}
			}
		}
	}
	println(answer)
}
