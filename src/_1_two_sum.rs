/*
 * @lc app=leetcode id=1 lang=rust
 *
 * [1] Two Sum
 */

// @lc code=start
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut m = std::collections::HashMap::with_capacity(nums.len());
        
        for (v, i) in nums.iter().zip(0..) {
            match m.get(&(target - *v)) {
                Some(&i2) => return vec![i2, i],
                None => {m.insert(*v, i);},
            }
        }
        panic!("none");
    }
}
// @lc code=end

