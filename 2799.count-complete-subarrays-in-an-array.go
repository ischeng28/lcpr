/*
 * @lc app=leetcode.cn id=2799 lang=golang
 * @lcpr version=30204
 *
 * [2799] 统计完全子数组的数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countCompleteSubarrays(nums []int) int {
	mp := map[int]struct{}{}
	for _, v := range nums {
		mp[v] = struct{}{}
	}
	diffCount := len(mp)

	count := map[int]int{}
	left := 0
	ans := 0
	for _, v := range nums {
		count[v]++
		for len(count) >= diffCount {
			count[nums[left]]--
			if count[nums[left]] == 0 {
				delete(count, nums[left])
			}
			left++
		}
		ans += left
	}
	return ans

}

// @lc code=end

/*
// @lcpr case=start
// [1,3,1,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [5,5,5,5]\n
// @lcpr case=end

*/

