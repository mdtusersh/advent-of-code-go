package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/mdtusersh/advent-of-code-go/util"
)

func main() {
	input := util.ReadLine("input.txt")

	fmt.Println("Part One:", partOne(input))
	fmt.Println("Part Two:", partTwo(input))
}

func partOne(input []string) int {
	ans := 0

	for _, s := range input {
		l, w, h := parseDimensions(s)

		dimensions := []int{l, w, h}
		sort.Ints(dimensions)

		area := (2 * l * w) + (2 * w * h) + (2 * h * l)
		extra := dimensions[0] * dimensions[1]

		ans += area + extra
	}

	return ans
}

func partTwo(input []string) int {
	ans := 0

	for _, s := range input {
		l, w, h := parseDimensions(s)

		dimensions := []int{l, w, h}
		sort.Ints(dimensions)

		present := l * w * h
		bow := (2 * dimensions[0]) + (2 * dimensions[1])

		ans += present + bow
	}

	return ans
}

func parseDimensions(s string) (int, int, int) {
	token := strings.Split(s, "x")

	l := util.MustAtoi(token[0])
	w := util.MustAtoi(token[1])
	h := util.MustAtoi(token[2])

	return l, w, h
}
