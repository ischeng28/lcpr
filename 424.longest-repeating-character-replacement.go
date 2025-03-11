/*
 * @lc app=leetcode.cn id=424 lang=golang
 * @lcpr version=30204
 *
 * [424] 替换后的最长重复字符
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func characterReplacement(s string, k int) int {
	// 创建一个 map 用于存储字符及其索引
	mp := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]] = append(mp[s[i]], i)
	}

	res := 0
	// 遍历 map 中的每个字符及其索引数组
	for _, arr := range mp {
		left := 0
		// 使用滑动窗口计算最大长度
		for right := 0; right < len(arr); right++ {
			for arr[right]-arr[left]+1-(right-left+1) > k {
				left++
			}
			res = max(res, right-left+1)
		}
	}

	// 确保结果不超过字符串长度
	if res+k > len(s) {
		return len(s)
	}
	return res + k
}

// @lc code=end

/*
// @lcpr case=start
// "ABAB"\n2\n
// @lcpr case=end

// @lcpr case=start
// "AABABBA"\n1\n
// @lcpr case=end

*/

