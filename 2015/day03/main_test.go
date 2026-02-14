package main

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := partOne(tt.input)
			if got != tt.want {
				t.Errorf("Part1(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}

}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := partTwo(tt.input)
			if got != tt.want {
				t.Errorf("Part1(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
