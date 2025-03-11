/*
 * @lc app=leetcode.cn id=1461 lang=golang
 * @lcpr version=30204
 *
 * [1461] 检查一个字符串是否包含所有长度为 K 的二进制子串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func hasAllCodes(s string, k int) bool {
	n := len(s)
	mp := map[string]struct{}{}
	for i := 0; i <= n-k; i++ {
		mp[s[i:i+k]] = struct{}{}
	}
	return len(mp) == 1<<k
}

// @lc code=end

/*
// @lcpr case=start
// "00110110"\n2\n
// @lcpr case=end

// @lcpr case=start
// "0110"\n1\n
// @lcpr case=end

// @lcpr case=start
// "0110"\n2\n
// @lcpr case=end

*/

