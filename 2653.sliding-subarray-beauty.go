/*
 * @lc app=leetcode.cn id=2653 lang=golang
 * @lcpr version=30204
 *
 * [2653] 滑动子数组的美丽值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func getSubarrayBeauty(nums []int, k int, x int) []int {
	ct := make([]int, 101)
	left := 0
	res := []int{}
	for right, v := range nums {
		ct[v+50]++
		if right-left+1 < k {
			continue
		}
		res = append(res, findXthNegative(ct, x))
		ct[nums[left]+50]--
		left++
	}
	return res
}

func findXthNegative(nums []int, x int) int {
	less := 0
	for i := 0; i < 50; i++ {
		less += nums[i]
		if less >= x {
			return i - 50
		}
	}
	return 0
}

// @lc code=end

/*
// @lcpr case=start
// [1,-1,-3,-2,3]\n3\n2\n
// @lcpr case=end

// @lcpr case=start
// [-1,-2,-3,-4,-5]\n2\n2\n
// @lcpr case=end

// @lcpr case=start
// [-3,1,2,-3,0,-3]\n2\n1\n
// @lcpr case=end

*/

