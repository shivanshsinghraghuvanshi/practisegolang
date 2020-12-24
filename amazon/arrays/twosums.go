package arrays

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	var output []int
	for i, num := range nums {
		m[num] = i
	}
	for i, num := range nums {
		c := target - num
		if value, ok := m[c]; ok && value != i {
			output = append(output, i, value)
			break
		}
	}
	return output
}
