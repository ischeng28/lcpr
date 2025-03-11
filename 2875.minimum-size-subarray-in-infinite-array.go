/*
 * @lc app=leetcode.cn id=2875 lang=golang
 * @lcpr version=30204
 *
 * [2875] 无限数组的最短子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minSizeSubarray(nums []int, target int) int {
	left := 0
	sum := 0
	n := len(nums)
	total := 0
	ans := math.MaxInt

	for _, v := range nums {
		total += v
	}
	for right := 0; right < 2*n; right++ {
		sum += nums[right%n]
		for sum >= target%total {
			if sum == target%total {
				ans = min(ans, right-left+1)
			}
			sum -= nums[left%n]
			left++
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans + target/total*n
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,2,3]\n4\n
// @lcpr case=end

// @lcpr case=start
// [2,4,6,8]\n3\n
// @lcpr case=end

*/

