/*
 * @lc app=leetcode.cn id=2958 lang=golang
 * @lcpr version=30204
 *
 * [2958] 最多 K 个重复元素的最长子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSubarrayLength(nums []int, k int) (ans int) {
	mp := map[int]int{}
	left := 0
	for right, _ := range nums {
		mp[nums[right]]++
		for mp[nums[right]] > k {
			mp[nums[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1,2,3,1,2]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1,2,1,2,1,2]\n1\n
// @lcpr case=end

// @lcpr case=start
// [5,5,5,5,5,5,5]\n4\n
// @lcpr case=end

*/

