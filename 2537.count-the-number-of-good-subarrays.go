/*
 * @lc app=leetcode.cn id=2537 lang=golang
 * @lcpr version=30204
 *
 * [2537] 统计好子数组的数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countGood(nums []int, k int) int64 {
	count := 0
	mp := map[int]int{}
	left := 0
	ans := 0
	for _, v := range nums {
		count += mp[v]
		mp[v]++
		for count >= k {
			mp[nums[left]]--
			count -= mp[nums[left]]
			left++
		}
		ans += left

	}
	return int64(ans)
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,1,1]\n10\n
// @lcpr case=end

// @lcpr case=start
// [3,1,4,3,2,2,4]\n2\n
// @lcpr case=end

*/

