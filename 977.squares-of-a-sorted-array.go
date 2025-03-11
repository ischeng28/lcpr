/*
 * @lc app=leetcode.cn id=977 lang=golang
 * @lcpr version=30204
 *
 * [977] 有序数组的平方
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	res := make([]int, len(nums))
	p := len(nums) - 1
	for left <= right {
		if abs(nums[left]) <= abs(nums[right]) {
			res[p] = nums[right] * nums[right]
			p--
			right--
		} else {
			res[p] = nums[left] * nums[left]
			p--
			left++
		}
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// @lc code=end

/*
// @lcpr case=start
// [-4,-1,0,3,10]\n
// @lcpr case=end

// @lcpr case=start
// [-7,-3,2,3,11]\n
// @lcpr case=end

*/

