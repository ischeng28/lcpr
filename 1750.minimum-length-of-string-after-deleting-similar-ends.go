/*
 * @lc app=leetcode.cn id=1750 lang=golang
 * @lcpr version=30204
 *
 * [1750] 删除字符串两端相同字符后的最短长度
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumLength(s string) int {
	n := len(s)
	left, right := 0, n-1
	for left < right && s[right] == s[left] {
		for left < right && s[right-1] == s[right] {
			right--
		}
		for left < right && s[left+1] == s[left] {
			left++
		}
		right--
		left++
	}

	return max(right-left+1, 0)
}

// @lc code=end

/*
// @lcpr case=start
// "ca"\n
// @lcpr case=end

// @lcpr case=start
// "cabaabac"\n
// @lcpr case=end

// @lcpr case=start
// "aabccabba"\n
// @lcpr case=end

*/

