/*
 * @lc app=leetcode.cn id=2260 lang=golang
 * @lcpr version=30204
 *
 * [2260] 必须拿起的最小连续卡牌数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumCardPickup(cards []int) int {
	res := math.MaxInt32
	mp := map[int]int{}
	for i, v := range cards {
		if e, ok := mp[v]; ok {
			res = min(res, i-e+1)
		}
		mp[v] = i
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [3,4,2,3,4,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,5,3]\n
// @lcpr case=end

*/

