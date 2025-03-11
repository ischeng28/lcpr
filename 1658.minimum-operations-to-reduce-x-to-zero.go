/*
 * @lc app=leetcode.cn id=1658 lang=golang
 * @lcpr version=30204
 *
 * [1658] 将 x 减到 0 的最小操作数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minOperations(nums []int, x int) (ans int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	target := sum - x
	if target < 0 {
		return -1
	}
	if target == 0 {
		return len(nums)
	}
	left := 0
	s := 0
	sublength := 0
	for right, _ := range nums {
		s += nums[right]
		for s > target {
			s -= nums[left]
			left++
		}
		if s == target {
			sublength = max(sublength, right-left+1)
		}
	}
	if sublength != 0 {
		return len(nums) - sublength
	}
	return -1

}

// @lc code=end

/*
// @lcpr case=start
// [1,1,4,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [5,6,7,8,9]\n4\n
// @lcpr case=end

// @lcpr case=start
// [3,2,20,1,1,3]\n10\n
// @lcpr case=end

*/

