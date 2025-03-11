/*
 * @lc app=leetcode.cn id=2486 lang=golang
 * @lcpr version=30204
 *
 * [2486] 追加字符以获得子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func appendCharacters(s string, t string) int {
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] != t[j] {
			continue
		} else {
			j++
		}
		if j == len(t) {
			return 0
		}
	}
	return len(t) - j
}

// @lc code=end

/*
// @lcpr case=start
// "coaching"\n"coding"\n
// @lcpr case=end

// @lcpr case=start
// "abcde"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "z"\n"abcde"\n
// @lcpr case=end

*/

