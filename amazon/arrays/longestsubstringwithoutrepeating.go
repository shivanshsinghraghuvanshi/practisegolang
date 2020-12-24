package arrays

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int, len(s))
	ans, b := 0, 0
	runes := []rune(s)
	if 0 == len(s) {
		return ans
	}
	for j, r := range runes {
		if _, ok := m[r]; ok {
			b = max(m[r], b)
		}
		ans = max(ans, j-b+1)
		m[r] = j + 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
