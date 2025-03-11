/*
 * @lc app=leetcode.cn id=713 lang=golang
 * @lcpr version=30204
 *
 * [713] 乘积小于 K 的子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	prod := 1
	left := 0
	ans := 0
	for right, _ := range nums {
		prod *= nums[right]
		for prod >= k {
			prod = prod / nums[left]
			left++
		}
		ans += right - left + 1
	}

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [10,5,2,6]\n100\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n0\n
// @lcpr case=end

*/

