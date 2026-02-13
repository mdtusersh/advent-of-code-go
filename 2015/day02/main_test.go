package main

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"2x3x4"}, 58},
		{[]string{"1x1x10"}, 43},
	}

	for _, tt := range tests {
		t.Run(tt.input[0], func(t *testing.T) {
			got := partOne(tt.input)
			if got != tt.want {
				t.Errorf("Part1(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}

}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"2x3x4"}, 34},
		{[]string{"1x1x10"}, 14},
	}

	for _, tt := range tests {
		t.Run(tt.input[0], func(t *testing.T) {
			got := partTwo(tt.input)
			if got != tt.want {
				t.Errorf("Part1(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
