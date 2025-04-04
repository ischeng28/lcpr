/*
 * @lc app=leetcode.cn id=1283 lang=golang
 * @lcpr version=30204
 *
 * [1283] 使结果不超过阈值的最小除数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func smallestDivisor(nums []int, threshold int) int {
	check := func(m int) bool {
		sum := 0
		for _, v := range nums {
			sum += (v + m - 1) / m
			if sum > threshold {
				return false
			}
		}
		return true
	}
	left, right := 0, slices.Max(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,5,9]\n6\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5,7,11]\n11\n
// @lcpr case=end

// @lcpr case=start
// [19]\n5\n
// @lcpr case=end

*/

