/*
 * @lc app=leetcode.cn id=2779 lang=golang
 * @lcpr version=30204
 *
 * [2779] 数组的最大美丽值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumBeauty(nums []int, k int) (ans int) {
	sort.Ints(nums)
	left := 0
	for right, _ := range nums {
		for nums[right]-k > nums[left]+k {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [4,6,1,2]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,1]\n10\n
// @lcpr case=end

*/

