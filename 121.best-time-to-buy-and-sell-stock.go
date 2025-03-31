/*
 * @lc app=leetcode.cn id=121 lang=golang
 * @lcpr version=30204
 *
 * [121] 买卖股票的最佳时机
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProfit(prices []int) int {
	min := math.MaxInt
	res := 0
	for _, v := range prices {
		if v < min {
			min = v
		} else {
			res = max(res, v-min)
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/

