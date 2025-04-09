/*
 * @lc app=leetcode.cn id=2559 lang=golang
 * @lcpr version=30204
 *
 * [2559] 统计范围内的元音字符串数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func vowelStrings(words []string, queries [][]int) []int {
	pre := make([]int, len(words)+1)
	for i, v := range words {
		if strings.Contains("aeiou", string(v[0])) && strings.Contains("aeiou", string(v[len(v)-1])) {
			pre[i+1] = pre[i] + 1
		} else {
			pre[i+1] = pre[i]
		}
	}

	res := make([]int, len(queries))
	for i, v := range queries {
		start, end := v[0], v[1]
		res[i] = pre[end+1] - pre[start]
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// ["aba","bcb","ece","aa","e"]\n[[0,2],[1,4],[1,1]]\n
// @lcpr case=end

// @lcpr case=start
// ["a","e","i"]\n[[0,2],[0,1],[2,2]]\n
// @lcpr case=end

*/

