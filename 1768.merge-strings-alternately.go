/*
 * @lc app=leetcode id=1768 lang=golang
 *
 * [1768] Merge Strings Alternately
 */

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	m := len(word1)
	n := len(word2)
	
	var result strings.Builder
	
	i,j := 0, 0
	
	for i < m || j < n {
		if i < m {
			result.WriteByte(word1[i])
			i++
		}
		if j < n {
			result.WriteByte(word2[j])
			j++
		}
	}
	
	return result.String()
}
// @lc code=end

