/*
 * @lc app=leetcode.cn id=643 lang=golang
 * @lcpr version=30204
 *
 * [643] 子数组最大平均数 I
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findMaxAverage(nums []int, k int) (res float64) {
	res = -math.MaxInt32
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i < k-1 {
			continue
		}
		res = max(res, float64(sum)/float64(k))
		sum -= nums[i-k+1]
	}
	return res

}

// @lc code=end

/*
// @lcpr case=start
// [1,12,-5,-6,50,3]\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n1\n
// @lcpr case=end

*/

