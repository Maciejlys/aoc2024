package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  140,
		},
		{
			name:  "example2",
			input: example2,
			want:  772,
		},
		{
			name:  "example3",
			input: example3,
			want:  1930,
		},
		{
			name:  "actual",
			input: actual,
			want:  1421958,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  80,
		},
		// {
		// 	name:  "example3",
		// 	input: example3,
		// 	want:  1206,
		// },
		// {
		// 	name:  "actual",
		// 	input: actual,
		// 	want:  885394,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
