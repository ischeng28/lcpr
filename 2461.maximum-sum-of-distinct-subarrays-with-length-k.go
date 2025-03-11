/*
 * @lc app=leetcode.cn id=2461 lang=golang
 * @lcpr version=30204
 *
 * [2461] 长度为 K 子数组中的最大和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	mp := map[int]int{}
	res := 0
	sum := 0
	start := 0
	for end := 0; end < n; end++ {
		sum += nums[end]
		mp[nums[end]]++
		if end-start+1 < k {
			continue
		}
		if len(mp) == k {
			res = max(res, sum)
		}

		sum -= nums[start]
		mp[nums[start]]--
		if mp[nums[start]] == 0 {
			delete(mp, nums[start])
		}

		start++
	}
	return int64(res)
}

// @lc code=end

/*
// @lcpr case=start
// [1,5,4,2,9,9,9]\n3\n
// @lcpr case=end

// @lcpr case=start
// [4,4,4]\n3\n
// @lcpr case=end

*/

