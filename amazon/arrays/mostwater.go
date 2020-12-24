package arrays

// Leetcode finding the max area of the container
func MaxArea(height []int) int {
	max := 0
	l := 0
	r := len(height) - 1
	//for i := 0; i < len(height); i++ {
	//	for j := 0; j < len(height); j++ {
	//		max = maxofTwo(max, minofTwo(height[i], height[j])*(j-i))
	//	}
	//}
	for l < r {
		max = maxofTwo(max, minofTwo(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

func maxofTwo(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func minofTwo(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
