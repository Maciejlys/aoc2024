package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		amount int
		size   int
		want   int
	}{
		{
			name:   "actual",
			input:  actual,
			size:   70,
			amount: 1024,
			want:   248,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, tt.size, tt.amount); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		size  int
		want  string
	}{
		{
			name:  "actual",
			input: actual,
			size:  70,
			want:  "32,55",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input, tt.size); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
