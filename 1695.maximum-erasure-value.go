/*
 * @lc app=leetcode.cn id=1695 lang=golang
 * @lcpr version=30204
 *
 * [1695] 删除子数组的最大得分
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumUniqueSubarray(nums []int) (ans int) {
	left := 0
	mp := map[int]int{}
	sum := 0
	for right, _ := range nums {
		mp[nums[right]]++
		sum += nums[right]
		for mp[nums[right]] > 1 {
			mp[nums[left]]--
			sum -= nums[left]
			left++
		}
		ans = max(ans, sum)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,4,5,6]\n
// @lcpr case=end

// @lcpr case=start
// [5,2,1,2,5,2,1,2,5]\n
// @lcpr case=end

*/

