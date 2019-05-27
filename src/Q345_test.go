package src

import "testing"

func Test_reverseVowels(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"leetcode",
			"leetcode",
			"leotcede",
		},
		{"aA",
			"aA",
			"Aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
