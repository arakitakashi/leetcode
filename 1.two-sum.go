/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if v, ok := m[complement]; ok {
			return []int{v, i}
		}
		m[num] = i
	}
	return []int{}
}

// @lc code=end

