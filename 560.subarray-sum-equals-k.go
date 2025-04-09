/*
 * @lc app=leetcode.cn id=560 lang=golang
 * @lcpr version=30204
 *
 * [560] 和为 K 的子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func subarraySum(nums []int, k int) int {
	mp := map[int]int{}
	sum := 0
	ans := 0
	for _, v := range nums {
		mp[sum]++
		sum += v
		ans += mp[sum-k]
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n3\n
// @lcpr case=end

*/

