/*
 * @lc app=leetcode.cn id=1021 lang=golang
 * @lcpr version=30204
 *
 * [1021] 删除最外层的括号
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func removeOuterParentheses(s string) string {
	stack := []byte{}
	var res string
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v == '(' {
			stack = append(stack, v)
			if len(stack) > 0 {
				res += string(v)
			}
		} else {
			if len(stack) > 1 {
				res += string(v)
			}
			stack = stack[:len(stack)-1]
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "(()())(())"\n
// @lcpr case=end

// @lcpr case=start
// "(()())(())(()(()))"\n
// @lcpr case=end

// @lcpr case=start
// "()()"\n
// @lcpr case=end

*/

