/*
 * @lc app=leetcode.cn id=2874 lang=golang
 * @lcpr version=30204
 *
 * [2874] 有序三元组中的最大值 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumTripletValue(nums []int) int64 {
	maxDiff, preMax, ans := 0, 0, 0
	for _, v := range nums {
		ans = max(ans, maxDiff*v)
		maxDiff = max(maxDiff, preMax-v)
		preMax = max(preMax, v)
	}
	return int64(ans)
}

// @lc code=end

/*
// @lcpr case=start
// [12,6,1,2,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,10,3,4,19]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/

