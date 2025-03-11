/*
 * @lc app=leetcode.cn id=2762 lang=golang
 * @lcpr version=30204
 *
 * [2762] 不间断子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func continuousSubarrays(nums []int) int64 {
	mp := map[int]int{}
	mn := math.MaxInt
	mx := math.MinInt
	left := 0
	ans := 0
	for right, v := range nums {
		mp[v]++
		mx = maxValue(mp)
		mn = minValue(mp)
		for mx > mn+2 {
			mp[nums[left]]--
			if mp[nums[left]] == 0 {
				delete(mp, nums[left])
			}
			mx = maxValue(mp)
			mn = minValue(mp)

			left++
		}
		ans += right - left + 1
	}
	return int64(ans)
}

func maxValue(mp map[int]int) int {
	mx := math.MinInt
	for k, _ := range mp {
		mx = max(k, mx)
	}
	return mx
}

func minValue(mp map[int]int) int {
	mn := math.MaxInt
	for k, _ := range mp {
		mn = min(k, mn)
	}
	return mn
}

// @lc code=end

/*
// @lcpr case=start
// [5,4,2,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/

