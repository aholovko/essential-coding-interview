// Find common elements in two sorted arrays.
package main

import "testing"

func TestCommonElem(t *testing.T) {
	tests := []struct {
		arg1 []int
		arg2 []int
		want []int
	}{
		{
			[]int{1, 2},
			[]int{2, 3},
			[]int{2},
		},
		{
			[]int{1, 3, 4, 6, 7, 9},
			[]int{1, 2, 4, 5, 9, 10},
			[]int{1, 4, 9},
		},
		{
			[]int{1, 2, 9, 10, 11, 12},
			[]int{0, 1, 2, 3, 4, 5, 8, 9, 10, 12, 14, 15},
			[]int{1, 2, 9, 10, 12},
		},
		{
			[]int{0, 1, 2, 3, 4, 5},
			[]int{6, 7, 8, 9, 10, 11},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := commonElem(tt.arg1, tt.arg2); !equals(got, tt.want) {
				t.Errorf("commonElem(%v, %v) = %v; want %v", tt.arg1, tt.arg2, got, tt.want)
			}
		})
	}
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	return true
}

func commonElem(a, b []int) []int {
	c := []int{}
	i, j := 0, 0
	for {
		if i >= len(a) || j >= len(b) {
			break
		}
		if a[i] == b[j] {
			c = append(c, a[i])
			i++
			j++
		} else if a[i] > b[j] {
			j++
		} else {
			i++
		}
	}
	return c
}
