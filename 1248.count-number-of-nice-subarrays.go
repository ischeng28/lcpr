/*
 * @lc app=leetcode.cn id=1248 lang=golang
 * @lcpr version=30204
 *
 * [1248] 统计「优美子数组」
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func numberOfSubarrays(nums []int, k int) int {
	count1, count2 := 0, 0
	left1, left2 := 0, 0
	ans := 0
	for _, v := range nums {
		if v%2 != 0 {
			count1++
			count2++
		}
		for count1 >= k {
			if nums[left1]%2 != 0 {
				count1--
			}
			left1++
		}
		for count2 >= k+1 {
			if nums[left2]%2 != 0 {
				count2--
			}
			left2++
		}
		ans += left1 - left2
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2,1,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,4,6]\n1\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,1,2,2,1,2,2,2]\n2\n
// @lcpr case=end

*/

