/*
 * @lc app=leetcode.cn id=2831 lang=golang
 * @lcpr version=30204
 *
 * [2831] 找出最长等值子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestEqualSubarray(nums []int, k int) (ans int) {
	maxSubCount := 0
	mp := map[int]int{}
	otherCount := 0
	left := 0
	for right, v := range nums {
		mp[v]++
		maxSubCount = max(maxSubCount, mp[v])
		otherCount = right - left + 1 - maxSubCount
		for otherCount > k {
			mp[nums[left]]--

		}
	}

}

// @lc code=end

/*
// @lcpr case=start
// [1,3,2,3,1,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,1,2,2,1,1]\n2\n
// @lcpr case=end

*/

