/*
 * @lc app=leetcode id=242 lang=rust
 *
 * [242] Valid Anagram
 */

// @lc code=start
impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        let mut exists = std::collections::HashMap::new();
        s.chars().for_each(|c| *exists.entry(c).or_insert(0) += 1);
        t.chars().for_each(|c| *exists.entry(c).or_insert(0) -= 1);
        exists.into_values().all(|v| v == 0)
    }
}
// @lc code=end

