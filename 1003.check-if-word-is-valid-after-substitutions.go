/*
 * @lc app=leetcode.cn id=1003 lang=golang
 * @lcpr version=30204
 *
 * [1003] 检查替换后的词是否有效
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isValid(s string) bool {
	stack := []rune{}
	for _, v := range s {
		if v == 'c' {
			if len(stack) < 2 {
				return false
			}
			if stack[len(stack)-1] == 'b' && stack[len(stack)-2] == 'a' {
				stack = stack[:len(stack)-2]
			} else {
				return false
			}
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0

}

// @lc code=end

/*
// @lcpr case=start
// "aabcbc"\n
// @lcpr case=end

// @lcpr case=start
// "abcabcababcc"\n
// @lcpr case=end

// @lcpr case=start
// "abccba"\n
// @lcpr case=end

*/

