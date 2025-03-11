/*
 * @lc app=leetcode.cn id=3185 lang=golang
 * @lcpr version=30204
 *
 * [3185] 构成整天的下标对数目 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countCompleteDayPairs(hours []int) int64 {
	mp := map[int]int{}
	res := 0
	for _, v := range hours {
		if e, ok := mp[(24-v%24)%24]; ok {
			res += e
		}
		mp[v%24]++
	}
	return int64(res)
}

// @lc code=end

/*
// @lcpr case=start
// [12,12,30,24,24]\n
// @lcpr case=end

// @lcpr case=start
// [72,48,24,3]\n
// @lcpr case=end

*/

