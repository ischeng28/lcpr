/*
 * @lc app=leetcode.cn id=2090 lang=golang
 * @lcpr version=30204
 *
 * [2090] 半径为 k 的子数组平均值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func getAverages(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		if i < k {
			res[i] = -1
		}
		if i >= n-k {
			res[i] = -1
		}
		sum += nums[i]
		if i < 2*k {
			continue
		}
		res[i-k] = sum / (2*k + 1)
		sum -= nums[i-2*k]
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [7,4,3,9,1,8,5,2,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [100000]\n0\n
// @lcpr case=end

// @lcpr case=start
// [8]\n100000\n
// @lcpr case=end

*/

