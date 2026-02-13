package main

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := part1(tt.input)
			if got != tt.want {
				t.Errorf("Part1(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}

}

func TestPart2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{")", 1},
		{"()())", 5},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := part2(tt.input)
			if got != tt.want {
				t.Errorf("solve2(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
