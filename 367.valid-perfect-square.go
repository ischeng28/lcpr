/*
 * @lc app=leetcode.cn id=367 lang=golang
 * @lcpr version=30204
 *
 * [367] 有效的完全平方数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isPerfectSquare(num int) bool {
	left, right := 0, num+1
	for left+1 < right {
		mid := left + (right-left)/2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			left = mid
		} else {
			right = mid
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// 16\n
// @lcpr case=end

// @lcpr case=start
// 14\n
// @lcpr case=end

*/

