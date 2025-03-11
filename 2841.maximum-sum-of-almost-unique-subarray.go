/*
 * @lc app=leetcode.cn id=2841 lang=golang
 * @lcpr version=30204
 *
 * [2841] 几乎唯一子数组的最大和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxSum(nums []int, m int, k int) int64 {
	n := len(nums)
	sum := 0
	mp := map[int]int{}
	ans := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		mp[nums[i]]++
		if i < k-1 {
			continue
		}
		if len(mp) >= m {
			ans = max(ans, sum)
		}
		sum -= nums[i-k+1]
		if mp[nums[i-k+1]] == 1 {
			delete(mp, nums[i-k+1])
		} else {
			mp[nums[i-k+1]]--
		}

	}
	return int64(ans)
}

// @lc code=end

/*
// @lcpr case=start
// [2,6,7,3,1,7]\n3\n4\n
// @lcpr case=end

// @lcpr case=start
// [5,9,9,2,4,5,4]\n1\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1,2,1,2,1]\n3\n3\n
// @lcpr case=end

*/

