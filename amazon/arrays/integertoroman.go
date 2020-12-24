package arrays

func IntToRoman(num int) string {
	var r []string
	var s string
	n := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 5, 4, 1}
	v := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "V", "IV", "I"}
	for i := 0; i < len(n); i++ {
		for n[i] <= num {
			num -= n[i]
			r = append(r, v[i])
		}
	}
	for _, x := range r {
		s += x
	}
	return s
}

func RomanToInt(s string) int {
	m := map[string]int{}
	m["M"] = 1000
	m["D"] = 500
	m["C"] = 100
	m["L"] = 50
	m["X"] = 10
	m["V"] = 5
	m["I"] = 1
	m["CM"] = 900
	m["CD"] = 400
	m["XC"] = 90
	m["XL"] = 40
	m["IX"] = 9
	m["IV"] = 4
	sum := 0
	i := 0
	for i < len(s) {
		if i < len(s)-1 {
			ss := s[i : i+2]
			if val, ok := m[ss]; ok {
				sum += val
				i += 2
				continue
			}
		}
		ss := s[i : i+1]
		sum += m[ss]
		i += 1
	}
	return sum
}

//M := []string{"", "M", "MM", "MMM"};
//C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"};
//X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"};
//I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"};
//
//return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
