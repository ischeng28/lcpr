/*
 * @lc app=leetcode.cn id=2904 lang=golang
 * @lcpr version=30204
 *
 * [2904] 最短且字典序最小的美丽子字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func shortestBeautifulSubstring(s string, k int) string {
	resStr := s
	left := 0
	n := len(s)
	count := 0
	if strings.Count(s, "1") < k {
		return ""
	}
	for right := 0; right < n; right++ {
		if s[right] == '1' {
			count++
		}
		for count >= k {
			str := s[left : right+1]
			if count == k && ((len(str) <= len(resStr) && str < resStr) || (len(str) < len(resStr))) {
				resStr = str
			}
			if s[left] == '1' {
				count--
			}
			left++
		}
	}
	return resStr
}

// @lc code=end

/*
// @lcpr case=start
// "100011001"\n3\n
// @lcpr case=end

// @lcpr case=start
// "1011"\n2\n
// @lcpr case=end

// @lcpr case=start
// "000"\n1\n
// @lcpr case=end

*/

