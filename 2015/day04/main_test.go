package main

import "testing"

func TestPartOne(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
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
