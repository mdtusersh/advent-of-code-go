package main

import (
	"crypto/md5"
	"fmt"
	"strconv"

	"github.com/mdtusersh/advent-of-code-go/util"
)

func main() {
	input := util.ReadInput("input.txt")

	fmt.Println("Part One:", partOne(input))
	fmt.Println("Part Two:", partTwo(input))
}

func partOne(input string) int {
	prefix := []byte(input)

	for i := 1; ; i++ {
		data := strconv.AppendInt(prefix, int64(i), 10)
		hash := md5.Sum(data)

		if hash[0] == 0 && hash[1] == 0 && hash[2] < 0x10 {
			return i
		}
	}
}

func partTwo(input string) int {
	prefix := []byte(input)

	for i := 1; ; i++ {
		data := strconv.AppendInt(prefix, int64(i), 10)
		hash := md5.Sum(data)

		if hash[0] == 0 && hash[1] == 0 && hash[2] == 0 {
			return i
		}
	}
}
