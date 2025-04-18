/*
 * @lc app=leetcode.cn id=3427 lang=golang
 * @lcpr version=30204
 *
 * [3427] 变长子数组求和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func subarraySum(nums []int) int {
	pre := make([]int, len(nums)+1)
	ans := 0
	for i, v := range nums {
		pre[i+1] = pre[i] + v
		if i-v > 0 {
			ans += pre[i+1] - pre[i-v]
		} else {
			ans += pre[i+1]
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [3,1,1,2]\n
// @lcpr case=end

*/

