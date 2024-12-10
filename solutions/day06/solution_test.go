package main

import (
	"image"
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
			want:  41,
		},
		{
			name:  "actual",
			input: actual,
			want:  4988,
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
			want:  6,
		},
		// {
		// 	name:  "actual",
		// 	input: actual,
		// 	want:  1697,
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

func Test_willGoOut(t *testing.T) {
	grid := map[image.Point]rune{
		{0, 0}: '^',
		{0, 1}: '.',
		{1, 0}: '#',
	}
	tests := []struct {
		orientation base
		want        bool
	}{
		{
			orientation: DOWN,
			want:        false,
		},
		{
			orientation: UP,
			want:        true,
		},
		{
			orientation: LEFT,
			want:        true,
		},
		{
			orientation: RIGHT,
			want:        false,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := willGoOut(tt.orientation, image.Point{0, 0}, grid); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
