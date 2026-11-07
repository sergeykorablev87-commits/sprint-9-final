package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{"zero size", 0, 0},
		{"negative size", -1, 0},
		{"one element", 1, 1},
		{"hundred elements", 100, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := generateRandomElements(tt.size)
			require.Len(t, res, tt.wantLen)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"positive numbers", []int{5, 1, 14, 1000, 365}, 1000},
		{"duplicates", []int{5, 1, 14, 1000, 365, 1000, 14, 1}, 1000},
		{"max is first", []int{2345, 5, 1, 14, 1000, 365}, 2345},
		{"max is last", []int{2345, 5, 1, 14, 1000, 365, 9999}, 9999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maximum(tt.slice)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
	}{
		{"short slice", []int{1, 7, 255, 9}},
		{"exactly chunks count", []int{1, 7, 255, 9, 134, 58, 245, 777}},
		{"correct result", []int{2345, 5, 1, 14, 1000, 365, 9999, 24626, 246246, 24624151, 6, 62626, 68906}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := maxChunks(tt.slice)
			require.Equal(t, maximum(tt.slice), res)
		})
	}
}
