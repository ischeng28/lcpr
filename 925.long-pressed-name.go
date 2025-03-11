/*
 * @lc app=leetcode.cn id=925 lang=golang
 * @lcpr version=30204
 *
 * [925] 长按键入
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for j < len(typed) {
		if i < len(name) && name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}
	return i == len(name)
}

// @lc code=end

/*
// @lcpr case=start
// "alex"\n"aaleex"\n
// @lcpr case=end

// @lcpr case=start
// "saeed"\n"ssaaedd"\n
// @lcpr case=end

*/

