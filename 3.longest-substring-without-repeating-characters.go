/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=30204
 *
 * [3] 无重复字符的最长子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func lengthOfLongestSubstring(s string) (ans int) {
	mp := map[byte]int{}
	left := 0
	for right, _ := range s {
		mp[s[right]]++
		for mp[s[right]] > 1 {
			mp[s[left]]--
			if mp[s[left]] == 0 {
				delete(mp, s[left])
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
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/

