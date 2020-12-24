package arrays

import "sort"

func ThreeSum(nums []int) [][]int {
	var o [][]int
	sort.Slice(nums, func(i, j int) bool {
		return i < j
	})
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		var res []int
		if i == 0 || nums[i-1] != nums[i] {
			res = TwoSum(nums, nums[i])
		}
		if len(res) == 2 {
			o = append(o, []int{nums[i], nums[res[0]], nums[res[1]]})
		}
	}
	return o
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

//func TwoSum(nums []int, target int) []int {
//	m := make(map[int]int, len(nums))
//	var output []int
//	for i, num := range nums {
//		m[num] = i
//	}
//	for i, num := range nums {
//		c := target - num
//		if value, ok := m[c]; ok && value != i {
//			output = append(output, i, value)
//			break
//		}
//	}
//	return output
//}
