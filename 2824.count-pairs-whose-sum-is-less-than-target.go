/*
 * @lc app=leetcode.cn id=2824 lang=golang
 * @lcpr version=30204
 *
 * [2824] 统计和小于目标的下标对数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	count := 0
	for left < right {
		tmp := nums[right] + nums[left]
		if tmp < target {
			count += right - left
			left++
		} else {
			right--
		}
	}
	return count
}

// @lc code=end

/*
// @lcpr case=start
// [-1,1,2,3,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [-6,2,5,-2,-7,-1,3]\n-2\n
// @lcpr case=end

*/

