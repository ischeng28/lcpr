/*
 * @lc app=leetcode.cn id=80 lang=golang
 * @lcpr version=30204
 *
 * [80] 删除有序数组中的重复项 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	left := 2
	for right := 2; right < len(nums); right++ {
		if nums[right] != nums[left-2] {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,2,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,1,1,2,3,3]\n
// @lcpr case=end

*/

