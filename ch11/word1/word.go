package word

func IsPalindrome(s string) bool {
	for i := range s {
		s1 := s[i]
		s2 := s[len(s)-i-1]
		if s1 != s2 {
			return false
		}
	}
	return true
}
