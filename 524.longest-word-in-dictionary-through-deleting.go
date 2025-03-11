/*
 * @lc app=leetcode.cn id=524 lang=golang
 * @lcpr version=30204
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findLongestWord(s string, dictionary []string) string {
	res := ""
	for _, e := range dictionary {
		i := 0
		for j := 0; j < len(s); j++ {
			if s[j] == e[i] {
				i++
				if i == len(e) {
					if len(e) > len(res) {
						res = e
					} else if len(e) == len(res) && e < res {
						res = e
					}
					break
				}
			}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// "abpcplea"\n["ale","apple","monkey","plea"]\n
// @lcpr case=end

// @lcpr case=start
// "abpcplea"\n["a","b","c"]\n
// @lcpr case=end

*/

