package stringutil

import (
	"testing"
	"fmt"
)

func TestRuneCount(t *testing.T) {
	fmt.Println("Testing RuneCount...")
	type testcase struct {
		in string 
		want map[rune]int
	}

	hellomap := make(map[rune]int)
	hellomap[rune('H')] = 1
	hellomap[rune('e')] = 1
	hellomap[rune('l')] = 2
	hellomap[rune('o')] = 1

	himap := make(map[rune]int)
	himap[rune('H')] = 1
	himap[rune('i')] = 1

	emptymap := make(map[rune]int)
	cases := []testcase{
		{"Hi", himap},
		{"Hello", hellomap},
		{"", emptymap},
	}

	for _, c := range cases {
		result := RuneCount(c.in)
		for key, _ := range result {
			if result[key] != c.want[key] {
				t.Errorf("Wrong!")
			}
		}
	}
}