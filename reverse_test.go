package stringutil

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct{
		in, want string
	}{
		{"Hello, wolrd", "drlow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for i, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
		fmt.Println("i", i)

	}
}