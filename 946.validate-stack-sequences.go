/*
 * @lc app=leetcode.cn id=946 lang=golang
 * @lcpr version=30204
 *
 * [946] 验证栈序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	j := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n[4,5,3,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n[4,3,5,1,2]\n
// @lcpr case=end

*/

