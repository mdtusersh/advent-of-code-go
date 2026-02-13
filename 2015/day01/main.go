package main

import (
	"fmt"

	"github.com/mdtusersh/advent-of-code-go/util"
)

func main() {
	input := util.ReadInput("input.txt")

	fmt.Println("Part One:", partOne(input))
	fmt.Println("Part Two:", partTwo(input))
}

func partOne(input string) int {
	ans := 0

	for _, value := range input {
		if value == '(' {
			ans++
		} else {
			ans--
		}
	}

	return ans
}

func partTwo(input string) int {
	ans := 0

	for index, value := range input {
		if value == '(' {
			ans++
		} else {
			ans--
		}

		if ans == -1 {
			return index + 1
		}
	}

	return 0
}
