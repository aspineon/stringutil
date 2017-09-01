package stringutil

func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	if s[0] == s[len(s)-1] {
		return IsPalindrome(s[1:len(s)-1])
	} else {
		return false
	}
}