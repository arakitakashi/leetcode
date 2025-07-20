/*
 * @lc app=leetcode id=217 lang=rust
 *
 * [217] Contains Duplicate
 */

// @lc code=start[]
impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut m = std::collections::HashSet::new();
        !nums.iter().all(|n| m.insert(n))
    }
}
// @lc code=end

