/*
 * @lc app=leetcode.cn id=1014 lang=golang
 * @lcpr version=30204
 *
 * [1014] 最佳观光组合
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxScoreSightseeingPair(values []int) (res int) {
	maxPre := 0
	for i, v := range values {
		res = max(res, maxPre+v-i)
		maxPre = max(maxPre, i+v)
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [8,1,5,2,6]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/

