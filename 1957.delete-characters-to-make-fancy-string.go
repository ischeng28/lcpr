/*
 * @lc app=leetcode.cn id=1957 lang=golang
 * @lcpr version=30204
 *
 * [1957] 删除字符使字符串变好
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func makeFancyString(s string) string {
	res := []rune{}
	i := 0
	for i < len(s) {
		start := i
		i++
		for i < len(s) && s[i] == s[i-1] {
			i++
		}
		length := min(2, i-start)
		for length > 0 {
			res = append(res, rune(s[start]))
			length--
		}
	}
	return string(res)
}

// @lc code=end

/*
// @lcpr case=start
// "leeetcode"\n
// @lcpr case=end

// @lcpr case=start
// "aaabaaaa"\n
// @lcpr case=end

// @lcpr case=start
// "aab"\n
// @lcpr case=end

*/

