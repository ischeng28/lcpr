/*
 * @lc app=leetcode.cn id=2825 lang=golang
 * @lcpr version=30204
 *
 * [2825] 循环增长使字符串子序列等于另一个字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func canMakeSubsequence(s string, t string) bool {
	if len(s) < len(t) {
		return false
	}
	j := 0
	for i := 0; i < len(s); i++ {
		c := s[i] + 1
		if s[i] == 'z' {
			c = 'a'
		}
		if t[j] == s[i] || t[j] == c {
			j++
			if j == len(t) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// "abc"\n"ad"\n
// @lcpr case=end

// @lcpr case=start
// "zc"\n"ad"\n
// @lcpr case=end

// @lcpr case=start
// "ab"\n"d"\n
// @lcpr case=end

*/

