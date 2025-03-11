/*
 * @lc app=leetcode.cn id=1234 lang=golang
 * @lcpr version=30204
 *
 * [1234] 替换子串得到平衡字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func balancedString(s string) int {
	mp := map[byte]int{}
	left := 0
	n := len(s)
	m := n / 4
	ans := n
	for i := 0; i < n; i++ {
		mp[s[i]]++
	}
	if mp['Q'] == m && mp['W'] == m && mp['E'] == m && mp['R'] == m {
		return 0
	}
	for right := 0; right < n; right++ {
		mp[s[right]]--
		for mp['Q'] <= m && mp['W'] <= m && mp['E'] <= m && mp['R'] <= m {
			ans = min(ans, right-left+1)
			mp[s[left]]++
			left++
		}
	}
	return ans

}

// @lc code=end

/*
// @lcpr case=start
// "QWER"\n
// @lcpr case=end

// @lcpr case=start
// "QQWE"\n
// @lcpr case=end

// @lcpr case=start
// "QQQW"\n
// @lcpr case=end

// @lcpr case=start
// "QQQQ"\n
// @lcpr case=end

*/

