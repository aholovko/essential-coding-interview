// Find where to expand in Minesweeper.
package main

import (
	"testing"
)

func TestMinesweeperExpand(t *testing.T) {
	tests := []struct {
		arg1 [][]int // board
		arg2 int     // rows
		arg3 int     // cols
		arg4 int     // click pos (x)
		arg5 int     // click pos (y)
		want [][]int // expanded board
	}{
		{
			[][]int{
				[]int{-1, 1, 0, 0},
				[]int{1, 1, 0, 0},
				[]int{0, 0, 1, 1},
				[]int{0, 0, 1, -1},
			},
			4,
			4,
			1,
			3,
			[][]int{
				[]int{-1, 1, -2, -2},
				[]int{1, 1, -2, -2},
				[]int{-2, -2, 1, 1},
				[]int{-2, -2, 1, -1},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minesweeperExpand(tt.arg1, tt.arg2, tt.arg3, tt.arg4, tt.arg5); !equals(got, tt.want) {
				t.Errorf("minesweeperExpand(%v, %d, %d, %d, %d) = %v, want %v", tt.arg1, tt.arg2, tt.arg3, tt.arg4, tt.arg5, got, tt.want)
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

func minesweeperExpand(board [][]int, rows, cols, x, y int) [][]int {
	if board[x][y] != 0 {
		return board
	}
	b := [][]int{}
	for i := 0; i < rows; i++ {
		b = append(b, make([]int, rows))
		for j := 0; j < cols; j++ {
			b[i][j] = board[i][j]
		}
	}
	q := [][]int{[]int{x, y}}
	for len(q) > 0 {
		xy := q[len(q)-1]
		q = q[:len(q)-1]
		x = xy[0]
		y = xy[1]
		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= rows {
				continue
			}
			for j := y - 1; j <= y+1; j++ {
				if j < 0 || j >= cols {
					continue
				}
				if b[i][j] == 0 {
					b[i][j] = -2
					q = append(q, []int{i, j})
				}
			}
		}
	}
	return b
}
