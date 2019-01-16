// Find if array A is a rotation of array B.
//
// Array is a rotation of another if it contains the same elements,
// which might or might not start at different index.
package main

import (
	"testing"
)

func TestIsRotation(t *testing.T) {
	tests := []struct {
		arg1 []int
		arg2 []int
		want bool
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{4, 5, 6, 7, 1, 2, 3},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 3, 4, 5, 6, 7},
			true,
		},
		{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			false,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{4, 5, 6, 7, 0, 2, 3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isRotation(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("isRotation(%v, %v) = %t, want %t", tt.arg1, tt.arg2, got, tt.want)
			}
			if got := isRotation2(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("isRotation2(%v, %v) = %t, want %t", tt.arg1, tt.arg2, got, tt.want)
			}
		})
	}
}

func isRotation(a, b []int) bool {
	check := false
	i, j := 0, 0
	for {
		if i >= len(a)-1 || j >= len(b)-1 {
			return check
		}
		if a[i] == b[j] {
			check = true
			i++
			j++
			if j >= len(b) {
				j = 0
			}
			continue
		}
		if check && a[i] != b[j] {
			return false
		}
		j++
	}
}

// alternate solution
func isRotation2(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	key := a[0]
	j := -1
	for i, v := range b {
		if v == key {
			j = i
			break
		}
	}
	if j == -1 {
		return false
	}
	lenb := len(b)
	for i, v := range a {
		if v != b[(j+i)%lenb] {
			return false
		}
	}
	return true
}
