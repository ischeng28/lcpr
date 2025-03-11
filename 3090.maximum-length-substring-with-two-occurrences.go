/*
 * @lc app=leetcode.cn id=3090 lang=golang
 * @lcpr version=30204
 *
 * [3090] 每个字符最多出现两次的最长子字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumLengthSubstring(s string) (ans int) {
	mp := map[byte]int{}
	left := 0
	for right, _ := range s {
		mp[s[right]]++
		for mp[s[right]] > 2 {
			mp[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans

}

// @lc code=end

/*
// @lcpr case=start
// "bcbbbcba"\n
// @lcpr case=end

// @lcpr case=start
// "aaaa"\n
// @lcpr case=end

*/

