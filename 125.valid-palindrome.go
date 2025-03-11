/*
 * @lc app=leetcode.cn id=125 lang=golang
 * @lcpr version=30204
 *
 * [125] 验证回文串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
			continue
		}
		if !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
			continue
		}
		if unicode.ToLower(rune(s[right])) == unicode.ToLower(rune(s[left])) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "A man, a plan, a canal: Panama"\n
// @lcpr case=end

// @lcpr case=start
// "race a car"\n
// @lcpr case=end

// @lcpr case=start
// " "\n
// @lcpr case=end

*/

