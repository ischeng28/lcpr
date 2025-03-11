/*
 * @lc app=leetcode.cn id=1438 lang=golang
 * @lcpr version=30204
 *
 * [1438] 绝对差不超过限制的最长连续子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestSubarray(nums []int, limit int) int {
	mp := map[int]int{}
	left := 0
	ans := 0
	mn := math.MaxInt
	mx := math.MinInt
	for right, v := range nums {
		mp[v]++
		mn = min(mn, v)
		mx = max(mx, v)
		for mx > mn+limit {
			mp[nums[left]]--
			if mp[nums[left]] == 0 {
				delete(mp, nums[left])
			}
			mn = minNum(mp)
			mx = maxNum(mp)
			left++
		}

		ans = max(ans, right-left+1)

	}
	return ans
}

func maxNum(mp map[int]int) int {
	mx := math.MinInt
	for key, _ := range mp {
		mx = max(mx, key)
	}
	return mx
}

func minNum(mp map[int]int) int {
	mn := math.MaxInt
	for key, _ := range mp {
		mn = min(mn, key)
	}
	return mn
}

// @lc code=end

/*
// @lcpr case=start
// [8,2,4,7]\n4\n
// @lcpr case=end

// @lcpr case=start
// [10,1,2,4,7,2]\n5\n
// @lcpr case=end

// @lcpr case=start
// [4,2,2,2,4,4,2,2]\n0\n
// @lcpr case=end

*/

