/*
 * @lc app=leetcode.cn id=2962 lang=golang
 * @lcpr version=30204
 *
 * [2962] 统计最大元素出现至少 K 次的子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countSubarrays(nums []int, k int) int64 {
	mx := slices.Max(nums)
	left := 0
	count := 0
	ans := 0
	for _, v := range nums {
		if v == mx {
			count++
		}
		for count >= k {
			if nums[left] == mx {
				count--
			}
			left++
		}
		ans += left
	}
	return int64(ans)

}

// @lc code=end

/*
// @lcpr case=start
// [1,3,2,3,3]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,4,2,1]\n3\n
// @lcpr case=end

*/

