/*
 * @lc app=leetcode.cn id=2024 lang=golang
 * @lcpr version=30204
 *
 * [2024] 考试的最大困扰度
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxConsecutiveAnswers(answerKey string, k int) (ans int) {
	return max(maxWithChar(answerKey, 'T', k), maxWithChar(answerKey, 'F', k))
}

func maxWithChar(key string, ch byte, k int) (ans int) {
	left := 0
	count := 0
	for right, _ := range key {
		if key[right] == ch {
			count++
		}
		for count > k {
			if key[left] == ch {
				count--
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
// "TTFF"\n2\n
// @lcpr case=end

// @lcpr case=start
// "TFFT"\n1\n
// @lcpr case=end

// @lcpr case=start
// "TTFTTFTT"\n1\n
// @lcpr case=end

*/

