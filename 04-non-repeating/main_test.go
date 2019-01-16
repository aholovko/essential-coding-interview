// Find non-repeating character in the string.
package main

import "testing"

func TestNonRepeatingChar(t *testing.T) {
	tests := []struct {
		arg  string
		want rune
	}{
		{"aabcb", 'c'},
		{"xxyz", 'y'},
		{"aabb", rune(0)},
		{"abab", rune(0)},
		{"aabbdbc", 'd'},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := nonRepeatingChar(tt.arg); got != tt.want {
				t.Errorf("nonRepeatingChar(%q) = %q, want %q", tt.arg, got, tt.want)
			}
		})
	}
}

func nonRepeatingChar(s string) rune {
	counts := map[rune]int{}
	for _, r := range s {
		counts[r]++
	}
	for k, v := range counts {
		if v == 1 {
			return k
		}
	}
	return rune(0)
}
