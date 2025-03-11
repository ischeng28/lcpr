/*
 * @lc app=leetcode.cn id=1869 lang=golang
 * @lcpr version=30204
 *
 * [1869] 哪种连续子字符串更长
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func checkZeroOnes(s string) bool {
	num1, num2 := 0, 0
	for i := 0; i < len(s); {
		start := i
		for ; i < len(s) && s[i] == s[start]; i++ {
		}
		if s[start] == '1' {
			num1 = max(num1, i-start)
		} else {
			num2 = max(num2, i-start)
		}
	}
	return num1 > num2
}

// @lc code=end

/*
// @lcpr case=start
// "1101"\n
// @lcpr case=end

// @lcpr case=start
// "111000"\n
// @lcpr case=end

// @lcpr case=start
// "110100010"\n
// @lcpr case=end

*/

