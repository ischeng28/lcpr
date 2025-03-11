/*
 * @lc app=leetcode.cn id=2529 lang=golang
 * @lcpr version=30204
 *
 * [2529] 正整数和负整数的最大计数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumCount(nums []int) int {
	nums1 := searchLowerBound(nums, 0)
	nums2 := searchLowerBound(nums, 1)
	return max(nums1, len(nums)-nums2)
}

func searchLowerBound(nums []int, target int) int {
	left, right := -1, len(nums)
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

// @lc code=end

/*
// @lcpr case=start
// [-2,-1,-1,1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [-3,-2,-1,0,0,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [5,20,66,1314]\n
// @lcpr case=end

*/

