// Rotate a 2D array by 90 degrees.
package main

import "testing"

func TestRotate(t *testing.T) {
	tests := []struct {
		arg  [][]int
		want [][]int
	}{
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{7, 4, 1},
				[]int{8, 5, 2},
				[]int{9, 6, 3},
			},
		},
		{
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
				[]int{13, 14, 15, 16},
			},
			[][]int{
				[]int{13, 9, 5, 1},
				[]int{14, 10, 6, 2},
				[]int{15, 11, 7, 3},
				[]int{16, 12, 8, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := rotate(tt.arg); !equals(got, tt.want) {
				t.Errorf("rotate(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func equals(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func copy(a [][]int) [][]int {
	b := [][]int{}
	for i := 0; i < len(a); i++ {
		b = append(b, make([]int, len(a)))
		for j := 0; j < len(a); j++ {
			b[i][j] = a[i][j]
		}
	}
	return b
}

func rotate(a [][]int) [][]int {
	b := copy(a)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			b[i][len(a)-j-1] = a[j][i]
		}
	}
	return b
}
