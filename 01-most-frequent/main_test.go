// Find the most frequent element in the array.
package main

import (
	"strconv"
	"testing"
)

func TestMostFreq(t *testing.T) {
	tests := []struct {
		arg  []int
		want int
	}{
		{nil, 0},
		{[]int{1}, 1},
		{[]int{1, 3, 1, 3, 2, 1}, 1},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := mostFrequent(tt.arg); got != tt.want {
				t.Errorf("mostFrequent(%v) = %d, want %d", tt.arg, got, tt.want)
			}
		})
	}
}

func mostFrequent(a []int) int {
	freq := make(map[int]int)
	maxCount, maxInt := -1, 0
	for _, e := range a {
		freq[e]++
		if freq[e] > maxCount {
			maxCount = freq[e]
			maxInt = e
		}
	}
	return maxInt
}
