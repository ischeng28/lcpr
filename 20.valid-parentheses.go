/*
 * @lc app=leetcode.cn id=20 lang=golang
 * @lcpr version=30204
 *
 * [20] 有效的括号
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isValid(s string) bool {
	stack := []byte{}
	mp := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < len(s); i++ {
		v := s[i]
		if pair, ok := mp[v]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == pair {
				stack = stack[:len(stack)-1]
				continue
			} else {
				stack = append(stack, v)
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
// "()"\n
// @lcpr case=end

// @lcpr case=start
// "()[]{}"\n
// @lcpr case=end

// @lcpr case=start
// "(]"\n
// @lcpr case=end

// @lcpr case=start
// "([])"\n
// @lcpr case=end

*/

