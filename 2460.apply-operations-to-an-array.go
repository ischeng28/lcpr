/*
 * @lc app=leetcode.cn id=2460 lang=golang
 * @lcpr version=30204
 *
 * [2460] 对数组执行操作
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func applyOperations(nums []int) []int {
	for i, _ := range nums[:len(nums)-1] {
		if nums[i] == nums[i+1] {
			nums[i] = nums[i] * 2
			nums[i+1] = 0
		}
	}
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
	}
	for left < len(nums) {
		nums[left] = 0
		left++
	}
	return nums
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,1,1,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,1]\n
// @lcpr case=end

*/

