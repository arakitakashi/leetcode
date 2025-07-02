/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}
	
	reversed := 0
	
	for x > reversed {
		reversed = reversed * 10 + x % 10
		x = x / 10
	}
	
	return x == reversed || reversed / 10 == x
}
// @lc code=end

