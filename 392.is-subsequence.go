/*
 * @lc app=leetcode.cn id=392 lang=golang
 * @lcpr version=30204
 *
 * [392] 判断子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	i := 0
	for j := 0; j < len(t); j++ {
		if t[j] != s[i] {
			continue
		} else {
			i++
		}
		if i == len(s) {
			return true
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// "abc"\n"ahbgdc"\n
// @lcpr case=end

// @lcpr case=start
// "axc"\n"ahbgdc"\n
// @lcpr case=end

*/

