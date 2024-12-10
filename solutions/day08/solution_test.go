package main

import (
	"image"
	"reflect"
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
			want:  14,
		},
		{
			name:  "actual",
			input: actual,
			want:  369,
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
			want:  34,
		},
		{
			name:  "actual",
			input: actual,
			want:  1169,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateAntinodesForPair(t *testing.T) {
	tests := []struct {
		a    image.Point
		b    image.Point
		want [2]image.Point
	}{
		{
			a:    image.Point{3, 4},
			b:    image.Point{5, 5},
			want: [2]image.Point{{1, 3}, {7, 6}},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := calculateAntinodesForPair(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
