/*
 * @lc app=leetcode.cn id=1456 lang=golang
 * @lcpr version=30204
 *
 * [1456] 定长子串中元音的最大数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxVowels(s string, k int) int {
	res, num := 0, 0
	n := len(s)
	start := 0
	for end := 0; end < n; end++ {
		if isVowels(rune(s[end])) {
			num = num + 1
		}
		if end-start+1 < k {
			continue
		}
		res = max(res, num)
		if isVowels(rune(s[start])) {
			num = num - 1
		}
		start++

	}
	return res
}

func isVowels(c rune) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// "abciiidef"\n3\n
// @lcpr case=end

// @lcpr case=start
// "aeiou"\n2\n
// @lcpr case=end

// @lcpr case=start
// "leetcode"\n3\n
// @lcpr case=end

// @lcpr case=start
// "rhythms"\n4\n
// @lcpr case=end

// @lcpr case=start
// "tryhard"\n4\n
// @lcpr case=end

*/

