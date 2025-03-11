/*
 * @lc app=leetcode.cn id=922 lang=golang
 * @lcpr version=30204
 *
 * [922] 按奇偶排序数组 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func sortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	n := len(nums)
	for odd < n && even < n {
		for even < n && nums[even]%2 == 0 {
			even += 2
		}
		for odd < n && nums[odd]%2 != 0 {
			odd += 2
		}
		if odd > n || even > n {
			return nums
		}
		nums[odd], nums[even] = nums[even], nums[odd]
		even += 2
		odd += 2
	}
	return nums

}

// @lc code=end

/*
// @lcpr case=start
// [4,2,5,7]\n
// @lcpr case=end

// @lcpr case=start
// [2,3]\n
// @lcpr case=end

*/

