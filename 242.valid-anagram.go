/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)
	n := make(map[rune]int)
	
	for i, u := range s {
		m[u]++
		n[rune(t[i])]++
	}
	
	for k, v := range m {
		if v != n[k] {
			return false
		}
	}
	
	return true
}
// @lc code=end

