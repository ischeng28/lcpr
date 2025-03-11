/*
 * @lc app=leetcode.cn id=633 lang=golang
 * @lcpr version=30204
 *
 * [633] 平方数之和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		v := left*left + right*right
		if v == c {
			return true
		} else if v < c {
			left++
		} else {
			right--
		}
	}
	return false

}

// @lc code=end

/*
// @lcpr case=start
// 5\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/

