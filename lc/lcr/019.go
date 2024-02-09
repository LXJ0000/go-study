package main

func validPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			// i ~ j - 1 || i + 1 ~ j
			if !check(s[i:j]) && !check(s[i+1:j+1]) {
				return false
			}
			return true
		}
	}
	return true
}

func check(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
