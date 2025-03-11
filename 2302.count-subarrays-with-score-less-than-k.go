/*
 * @lc app=leetcode.cn id=2302 lang=golang
 * @lcpr version=30204
 *
 * [2302] 统计得分小于 K 的子数组数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countSubarrays(nums []int, k int64) int64 {
	sum := 0
	left := 0
	ans := 0
	for right, v := range nums {
		sum += v
		for sum*(right-left+1) >= int(k) {
			sum -= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return int64(ans)
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,4,3,5]\n10\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1]\n5\n
// @lcpr case=end

*/

