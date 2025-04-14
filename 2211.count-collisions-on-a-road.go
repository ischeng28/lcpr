/*
 * @lc app=leetcode.cn id=2211 lang=golang
 * @lcpr version=30204
 *
 * [2211] 统计道路上的碰撞次数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countCollisions(s string) int {
	s = strings.TrimLeft(s, "L")
	s = strings.TrimRight(s, "R")
	return len(s) - strings.Count(s, "S")
}

// @lc code=end

/*
// @lcpr case=start
// "RLRSLL"\n
// @lcpr case=end

// @lcpr case=start
// "LLRR"\n
// @lcpr case=end

*/

