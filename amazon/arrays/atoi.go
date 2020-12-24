package arrays

import (
	"math"
	"strconv"
	"strings"
)

// MyAtoi is the code
func MyAtoi(s string) int {
	pos := 0
	s = strings.TrimSpace(s)
	if len(s) == 0 || (len(s) == 1 && (s[0] == '+' || s[0] == '-')) {
		return 0
	}
	if s[0] == '+' || s[0] == '-' {
		if _, err := strconv.Atoi(string(s[1])); err != nil {
			return 0
		}
		for p, r := range s[1:] {
			if _, err := strconv.Atoi(string(r)); err != nil {
				break
			}
			pos = p + 1
		}
	} else {
		if _, err := strconv.Atoi(string(s[0])); err != nil {
			return 0
		}
		for p, r := range s {
			if _, err := strconv.Atoi(string(r)); err != nil {
				break
			}
			pos = p
		}

	}
	x, err := strconv.ParseInt(s[:pos+1], 10, 32)
	if err != nil {
		if s[0] == '-' {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}
	return int(x)

}
