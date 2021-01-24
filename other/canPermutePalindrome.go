package other

func canPermutePalindrome(s string) bool {
	even := false
	check := true
	f := make(map[rune]int)
	if len(s)%2 == 0 {
		even = true
	}
	if even {
		for _, r := range s {
			if val, ok := f[r]; ok {
				f[r] = val + 1
			}
			f[r] = 1
		}
		for _, v := range f {
			m := v % 2
			if m != 0 {
				check = false
			}
		}
		return check
	} else {
		return check
	}
}
