/*
 * @lc app=leetcode.cn id=344 lang=golang
 * @lcpr version=30204
 *
 * [344] 反转字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func reverseString(s []byte) {
	var temp byte
	left, right := 0, len(s)-1
	for left < right {
		temp = s[right]
		s[right] = s[left]
		s[left] = temp
		left++
		right--
	}
}

// @lc code=end

/*
// @lcpr case=start
// ["h","e","l","l","o"]\n
// @lcpr case=end

// @lcpr case=start
// ["H","a","n","n","a","h"]\n
// @lcpr case=end

*/

