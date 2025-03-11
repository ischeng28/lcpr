/*
 * @lc app=leetcode.cn id=1446 lang=golang
 * @lcpr version=30204
 *
 * [1446] 连续字符
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxPower(s string) int {
	cnt, res := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
			if cnt > res {
				res = cnt
			}
		} else {
			cnt = 1
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "leetcode"\n
// @lcpr case=end

// @lcpr case=start
// "abbcccddddeeeeedcba"\n
// @lcpr case=end

*/

