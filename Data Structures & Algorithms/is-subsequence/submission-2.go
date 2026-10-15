func isSubsequence(s string, t string) bool {
	sptr, tptr := 0, 0

	for sptr < len(s) && tptr < len(t) {
		if s[sptr] == t[tptr] {
			sptr += 1
		}
		tptr += 1
	}
	return sptr == len(s)
}
