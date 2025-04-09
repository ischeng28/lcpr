/*
 * @lc app=leetcode.cn id=930 lang=golang
 * @lcpr version=30204
 *
 * [930] 和相同的二元子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numSubarraysWithSum(nums []int, goal int) (ans int) {
	left1, left2 := 0, 0
	sum1, sum2 := 0, 0
	for right, v := range nums {
		sum1 += v
		sum2 += v
		for left1 <= right && sum1 > goal {
			sum1 -= nums[left1]
			left1++
		}
		for left2 <= right && sum2 >= goal {
			sum2 -= nums[left2]
			left2++
		}
		ans += left2 - left1
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,0,1,0,1]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0,0,0]\n0\n
// @lcpr case=end

*/

