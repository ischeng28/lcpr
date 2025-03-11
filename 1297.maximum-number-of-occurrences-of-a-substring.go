/*
 * @lc app=leetcode.cn id=1297 lang=golang
 * @lcpr version=30204
 *
 * [1297] 子串的最大出现次数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	mp := map[rune]int{}
	count := map[string]int{}
	start := 0
	n := len(s)
	for end := 0; end < n; end++ {
		mp[rune(s[end])]++
		if end-start+1 < minSize {
			continue
		}
		if len(mp) <= maxLetters {
			count[s[start:end+1]]++
		}
		mp[rune(s[start])]--
		if mp[rune(s[start])] == 0 {
			delete(mp, rune(s[start]))
		}
		start++
	}
	ans := 0
	for _, v := range count {
		ans = max(ans, v)
	}
	return ans

}

// @lc code=end

/*
// @lcpr case=start
// "aababcaab"\n2\n3\n4\n
// @lcpr case=end

// @lcpr case=start
// "aaaa"\n1\n3\n3\n
// @lcpr case=end

// @lcpr case=start
// "aabcabcab"\n2\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// "abcde"\n2\n3\n3\n
// @lcpr case=end

*/

