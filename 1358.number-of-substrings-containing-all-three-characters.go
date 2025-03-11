/*
 * @lc app=leetcode.cn id=1358 lang=golang
 * @lcpr version=30204
 *
 * [1358] 包含所有三种字符的子字符串数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numberOfSubstrings(s string) int {
	mp := map[byte]int{}
	left := 0
	n := len(s)
	ans := 0
	for right := 0; right < n; right++ {
		mp[s[right]]++
		for mp['a'] >= 1 && mp['b'] >= 1 && mp['c'] >= 1 {
			mp[s[left]]--
			left++
		}
		ans += left
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// "abcabc"\n
// @lcpr case=end

// @lcpr case=start
// "aaacb"\n
// @lcpr case=end

// @lcpr case=start
// "abc"\n
// @lcpr case=end

*/

