/*
 * @lc app=leetcode.cn id=611 lang=golang
 * @lcpr version=30204
 *
 * [611] 有效三角形的个数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	n := len(nums)
	for right := n - 1; right >= 2; right-- {
		left, mid := 0, right-1
		for left < mid {
			if nums[left]+nums[mid] > nums[right] {
				res += mid - left
				mid--
			} else {
				left++
			}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [2,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [4,2,3,4]\n
// @lcpr case=end

*/

