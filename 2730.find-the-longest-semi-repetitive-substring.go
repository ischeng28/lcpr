/*
 * @lc app=leetcode.cn id=2730 lang=golang
 * @lcpr version=30204
 *
 * [2730] 找到最长的半重复子字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestSemiRepetitiveSubstring(s string) (ans int) {
	count := 0
	left := 0
	for right, _ := range s {
		if right > 0 && s[right] == s[right-1] {
			count++
		}
		for count > 1 {
			if s[left] == s[left+1] {
				count--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "52233"\n
// @lcpr case=end

// @lcpr case=start
// "5494"\n
// @lcpr case=end

// @lcpr case=start
// "1111111"\n
// @lcpr case=end

*/

