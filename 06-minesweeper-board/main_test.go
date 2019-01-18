// Build board for Minesweeper (assign numbers).
package main

import (
	"testing"
)

func TestMinesweeperBoard(t *testing.T) {
	tests := []struct {
		arg1 [][]int // mines positions
		arg2 int     // rows
		arg3 int     // columns
		want [][]int // minesweeper board with assigned numbers
	}{
		{
			[][]int{
				[]int{0, 0},
				[]int{0, 1},
			},
			3,
			4,
			[][]int{
				[]int{-1, -1, 1, 0},
				[]int{2, 2, 1, 0},
				[]int{0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minesweeperBoard(tt.arg1, tt.arg2, tt.arg3); !equals(got, tt.want) {
				t.Errorf("minesweeperBoard(%v, %d, %d) = %v, want %v", tt.arg1, tt.arg2, tt.arg3, got, tt.want)
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

func minesweeperBoard(mines [][]int, rows, cols int) [][]int {
	board := [][]int{}
	for i := 0; i < rows; i++ {
		board = append(board, make([]int, cols))
	}
	for _, xy := range mines {
		x := xy[0]
		y := xy[1]
		board[x][y] = -1
		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= rows {
				continue
			}
			for j := y - 1; j <= y+1; j++ {
				if j < 0 || j >= cols {
					continue
				}
				if board[i][j] != -1 {
					board[i][j]++
				}
			}
		}
	}
	return board
}
