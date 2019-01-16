// Find if strings are one away strings.
//
// One away strings differ at most in one character.
package main

import (
	"testing"
)

func TestIsOneAway(t *testing.T) {
	tests := []struct {
		arg1 string
		arg2 string
		want bool
	}{
		{"abcde", "abcd", true},
		{"abde", "abcde", true},
		{"a", "a", true},
		{"abcdef", "abqdef", true},
		{"abcdef", "abccef", true},
		{"abcdef", "abcde", true},
		{"aaa", "abc", false},
		{"abcde", "abc", false},
		{"abc", "abcde", false},
		{"abc", "bcc", false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isOneAway(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("isOneAway(%q, %q) = %t, want %t", tt.arg1, tt.arg2, got, tt.want)
			}
		})
	}
}

func isOneAway(s1, s2 string) bool {
	if abs(len(s1)-len(s2)) > 1 {
		return false
	}
	diff := 0
	if len(s1) == len(s2) {
		for i, c := range s1 {
			if c != rune(s2[i]) {
				diff++
				if diff > 1 {
					return false
				}
			}
		}
		return true
	}
	if len(s1) > len(s2) {
		return compareDiff(s1, s2)
	}
	return compareDiff(s2, s1)
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

// s1 is one character longer than s2
func compareDiff(s1, s2 string) bool {
	diff := 0
	for i := 0; i < len(s2); {
		if s1[i+diff] != s2[i] {
			diff++
			if diff > 1 {
				return false
			}
			continue
		}
		i++
	}
	return true
}
