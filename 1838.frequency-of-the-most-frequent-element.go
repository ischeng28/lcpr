/*
 * @lc app=leetcode.cn id=1838 lang=golang
 * @lcpr version=30204
 *
 * [1838] 最高频元素的频数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxFrequency(nums []int, k int) (ans int) {
	sort.Ints(nums)
	if len(nums) == 1 {
		return 1
	}
	sum := 0
	left := 0
	for right := 1; right < len(nums); right++ {
		sum += (right - left) * (nums[right] - nums[right-1])
		if sum > k {
			sum -= nums[right] - nums[left]
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,4]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1,4,8,13]\n5\n
// @lcpr case=end

// @lcpr case=start
// [3,9,6]\n2\n
// @lcpr case=end

*/

