/*
 * @lc app=leetcode.cn id=668 lang=golang
 * @lcpr version=30204
 *
 * [668] 乘法表中第k小的数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findKthNumber(m int, n int, k int) int {
	i, j := 0, n
	for i <= m-1 && j >= 0 {
		v := (i + 1) * (n + 1)
		if v {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

// @lc code=end

/*
// @lcpr case=start
// 3\n3\n5\n
// @lcpr case=end

// @lcpr case=start
// 2\n3\n6\n
// @lcpr case=end

*/

