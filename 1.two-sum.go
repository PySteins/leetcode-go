package leetcode

// https://leetcode-cn.com/problems/two-sum/
// 两数之和
func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i, num := range nums {
		if v, ok := cache[target-num]; ok {
			return []int{v, i}
		} else {
			cache[num] = i
		}
	}
	return []int{}
}
