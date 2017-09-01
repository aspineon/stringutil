package stringutil

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// True Cases
	fmt.Println("Testing IsPalindrome...")
	racecar := "racecar"
	empty := ""
	parens := "(()()(("
	single := "1"

	// False Cases
	trivial := "1234567789"
	lastTwo := "123421"

	type ipTestCase struct {
		in string
		want bool
	}
	cases := []ipTestCase{
		{racecar, true},
		{empty, true},
		{parens, true},
		{single, true},

		{trivial, false},
		{lastTwo, false},
	}
	for _, c := range cases{
		result := IsPalindrome(c.in)
		if result != c.want {
			t.Errorf("Wrong")
		}
	}


}