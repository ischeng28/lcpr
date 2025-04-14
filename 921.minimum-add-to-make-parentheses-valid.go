/*
 * @lc app=leetcode.cn id=921 lang=golang
 * @lcpr version=30204
 *
 * [921] 使括号有效的最少添加
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minAddToMakeValid(s string) int {
	stack := []byte{}
	count := 0
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v == '(' {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 {
			count++
			continue
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	count += len(stack)
	return count
}

// @lc code=end

/*
// @lcpr case=start
// "())"\n
// @lcpr case=end

// @lcpr case=start
// "((("\n
// @lcpr case=end

*/

