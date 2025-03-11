/*
 * @lc app=leetcode.cn id=441 lang=golang
 * @lcpr version=30204
 *
 * [441] 排列硬币
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func arrangeCoins(n int) int {
	left, right := 0, n+1
	for left+1 < right {
		mid := left + (right-left)/2
		if mid*(mid+1) <= 2*n {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}

// @lc code=end

/*
// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 8\n
// @lcpr case=end

*/

