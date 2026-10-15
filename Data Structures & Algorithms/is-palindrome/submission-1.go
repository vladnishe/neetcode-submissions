func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for l < r {
		if !isAlphaNum(s[l]) {
			l++
			continue
		} else if !isAlphaNum(s[r]) {
			r--
			continue
		}

		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') 
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}