/*
 * @lc app=leetcode.cn id=2390 lang=golang
 * @lcpr version=30204
 *
 * [2390] 从字符串中移除星号
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func removeStars(s string) string {
	stack := []rune{}
	for _, v := range s {
		if v == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}

// @lc code=end

/*
// @lcpr case=start
// "leet**cod*e"\n
// @lcpr case=end

// @lcpr case=start
// "erase*****"\n
// @lcpr case=end

*/

