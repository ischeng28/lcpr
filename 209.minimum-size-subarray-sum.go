/*
 * @lc app=leetcode.cn id=209 lang=golang
 * @lcpr version=30204
 *
 * [209] 长度最小的子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minSubArrayLen(target int, nums []int) (ans int) {
	left := 0
	ans = len(nums) + 1
	sum := 0
	for right, v := range nums {
		sum += v
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// 7\n[2,3,1,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[1,4,4]\n
// @lcpr case=end

// @lcpr case=start
// 11\n[1,1,1,1,1,1,1,1]\n
// @lcpr case=end

*/

