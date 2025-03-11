/*
 * @lc app=leetcode.cn id=1208 lang=golang
 * @lcpr version=30204
 *
 * [1208] 尽可能使字符串相等
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func equalSubstring(s string, t string, maxCost int) (ans int) {
	subSum := 0
	left := 0
	for right, _ := range s {
		subSum += absSub(s[right], t[right])
		for subSum > maxCost {
			subSum -= absSub(s[left], t[left])
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func absSub(a, b byte) int {
	if a > b {
		return int(a - b)
	} else {
		return int(b - a)
	}
}

// @lc code=end

/*
// @lcpr case=start
// "abcd"\n"bcdf"\n3\n
// @lcpr case=end

// @lcpr case=start
// "abcd"\n"cdef"\n3\n
// @lcpr case=end

// @lcpr case=start
// "abcd"\n"acde"\n0\n
// @lcpr case=end

*/

