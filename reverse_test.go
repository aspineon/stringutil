package stringutil

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Println("Testing Reverse...")
	cases := []struct{
		in, want string
	}{
		{"Hello, wolrd", "drlow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}