/*
 * @lc app=leetcode.cn id=844 lang=golang
 * @lcpr version=30204
 *
 * [844] 比较含退格的字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func backspaceCompare(s string, t string) bool {
	stack1 := []rune{}
	stack2 := []rune{}
	for _, v := range s {
		if v == '#' {
			if len(stack1) > 0 {
				stack1 = stack1[:len(stack1)-1]
			}
		} else {
			stack1 = append(stack1, v)
		}
	}
	for _, v := range t {
		if v == '#' {
			if len(stack2) > 0 {
				stack2 = stack2[:len(stack2)-1]
			}

		} else {
			stack2 = append(stack2, v)
		}
	}
	if len(stack1) != len(stack2) {
		return false
	}
	for i := 0; i < len(stack1); i++ {
		if stack1[i] != stack2[i] {
			return false
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "ab#c"\n"ad#c"\n
// @lcpr case=end

// @lcpr case=start
// "ab##"\n"c#d#"\n
// @lcpr case=end

// @lcpr case=start
// "a#c"\n"b"\n
// @lcpr case=end

*/

