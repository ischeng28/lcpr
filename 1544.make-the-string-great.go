/*
 * @lc app=leetcode.cn id=1544 lang=golang
 * @lcpr version=30204
 *
 * [1544] 整理字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func makeGood(s string) string {
	stack := []rune{}
	for _, v := range s {
		if len(stack) > 0 && (v-stack[len(stack)-1] == 32 || stack[len(stack)-1]-v == 32) {
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
// "leEeetcode"\n
// @lcpr case=end

// @lcpr case=start
// "abBAcC"\n
// @lcpr case=end

// @lcpr case=start
// "s"\n
// @lcpr case=end

*/

