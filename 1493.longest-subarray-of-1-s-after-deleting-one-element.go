/*
 * @lc app=leetcode.cn id=1493 lang=golang
 * @lcpr version=30204
 *
 * [1493] 删掉一个元素以后全为 1 的最长子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestSubarray(nums []int) (ans int) {
	left := 0
	zeroCount := 0
	for right, v := range nums {
		if v == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		ans = max(ans, right-left)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1,1,0,1,1,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1]\n
// @lcpr case=end

*/

