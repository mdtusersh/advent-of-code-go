package util

import (
	"os"
	"strconv"
	"strings"
)

func ReadInput(path string) string {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(data))
}

func ReadLine(path string) []string {
	return strings.Split(ReadInput(path), "\n")
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return n
}
