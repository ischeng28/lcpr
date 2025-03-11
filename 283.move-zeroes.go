/*
 * @lc app=leetcode.cn id=283 lang=golang
 * @lcpr version=30204
 *
 * [283] 移动零
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func moveZeroes(nums []int) {
	left := 0
	for right, v := range nums {
		if v != 0 {
			nums[left] = nums[right]
			left++
		}
	}
	for ; left < len(nums); left++ {
		nums[left] = 0
	}
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,0,3,12]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

