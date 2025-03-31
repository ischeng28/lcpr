/*
 * @lc app=leetcode.cn id=2576 lang=golang
 * @lcpr version=30204
 *
 * [2576] 求出最多标记下标
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxNumOfMarkedIndices(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	res := 0
	slices.Sort(nums)
	left, right := 0, 1
	for right < len(nums) {
		if nums[right] >= 2*nums[left] {
			res += 2
			left++
			right++
		} else {
			right++
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [3,5,2,4]\n
// @lcpr case=end

// @lcpr case=start
// [9,2,5,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,8]\n
// @lcpr case=end

*/

