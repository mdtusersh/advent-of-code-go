package main

import (
	"fmt"

	"github.com/mdtusersh/advent-of-code-go/util"
)

type Pair struct {
	A int
	B int
}

func main() {
	input := util.ReadInput("input.txt")

	fmt.Println("Part One:", partOne(input))
	fmt.Println("Part Two:", partTwo(input))
}

func partOne(input string) int {
	unique := make(map[Pair]struct{})

	pair := Pair{0, 0}

	unique[pair] = struct{}{}

	for _, direction := range input {
		move(&pair, direction)
		unique[pair] = struct{}{}
	}

	return len(unique)
}

func partTwo(input string) int {
	unique := make(map[Pair]struct{})

	pairSanta := Pair{0, 0}
	pairRobo := Pair{0, 0}

	unique[Pair{0, 0}] = struct{}{}

	for index, direction := range input {
		if index%2 == 0 {
			move(&pairSanta, direction)
			unique[pairSanta] = struct{}{}
		} else {
			move(&pairRobo, direction)
			unique[pairRobo] = struct{}{}
		}
	}

	return len(unique)
}

func move(pair *Pair, direction rune) {
	switch direction {
	case '^':
		pair.B++
	case 'v':
		pair.B--
	case '<':
		pair.A--
	case '>':
		pair.A++
	}
}
