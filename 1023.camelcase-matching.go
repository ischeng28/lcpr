/*
 * @lc app=leetcode.cn id=1023 lang=golang
 * @lcpr version=30204
 *
 * [1023] 驼峰式匹配
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func camelMatch(queries []string, pattern string) []bool {
	n := len(queries)
	ans := make([]bool, n)

	for i, query := range queries {
		ans[i] = true
		j := 0
		for _, q := range query {
			// 两个置为 false 的条件
			// 1、已经匹配到全部字符了但仍然有大写字母的
			// 2、还没有匹配完全部字符但是有不属于 pattern 的大写字母
			if (j == len(pattern) && unicode.IsUpper(q)) || (j < len(pattern) && unicode.IsUpper(q) && q != rune(pattern[j])) {
				ans[i] = false
				break
			}
			// 如果匹配到了则移动指针
			if j < len(pattern) && q == rune(pattern[j]) {
				j++
			}
		}
		// 循环完没有匹配完的也需要置为 false
		if j < len(pattern) {
			ans[i] = false
		}
	}

	return ans
}

// @lc code=end

/*
// @lcpr case=start
// ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"]\n"FB"\n
// @lcpr case=end

// @lcpr case=start
// ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"]\n"FoBa"\n
// @lcpr case=end

// @lcpr case=start
// ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"]\n"FoBaT"\n
// @lcpr case=end

*/

