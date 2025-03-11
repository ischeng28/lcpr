/*
 * @lc app=leetcode.cn id=69 lang=golang
 * @lcpr version=30204
 *
 * [69] x 的平方根
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func mySqrt(x int) int {
	left, right := 0, x+1
	for left+1 < right {
		mid := left + (right-left)/2
		if mid*mid <= x {
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
// 4\n
// @lcpr case=end

// @lcpr case=start
// 8\n
// @lcpr case=end

*/

