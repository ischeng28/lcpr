/*
 * @lc app=leetcode.cn id=905 lang=golang
 * @lcpr version=30204
 *
 * [905] 按奇偶排序数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func sortArrayByParity(nums []int) []int {
	left := 0
	for right, v := range nums {
		if v%2 == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
	return nums
}

// @lc code=end

/*
// @lcpr case=start
// [3,1,2,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

