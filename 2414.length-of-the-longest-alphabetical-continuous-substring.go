/*
 * @lc app=leetcode.cn id=2414 lang=golang
 * @lcpr version=30204
 *
 * [2414] 最长的字母序连续子字符串的长度
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestContinuousSubstring(s string) int {
	i := 0
	res := 0
	for i < len(s) {
		start := i
		i = i + 1
		for ; i < len(s) && s[i] == s[i-1]+1; i++ {
		}
		res = max(res, i-start)
	}
	return res
}

func longestContinuousSubstring(s string) int {
	ans, cnt := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i-1]+1 == s[i] {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt = 1
		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "abacaba"\n
// @lcpr case=end

// @lcpr case=start
// "abcde"\n
// @lcpr case=end

*/

