/*
 * @lc app=leetcode.cn id=1004 lang=golang
 * @lcpr version=30204
 *
 * [1004] 最大连续1的个数 III
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func longestOnes(nums []int, k int) (ans int) {
	left := 0
	count := 0
	for right, _ := range nums {
		if nums[right] == 0 {
			count++
		}
		for count > k {
			if nums[left] == 0 {
				count--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,0,0,0,1,1,1,1,0]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1]\n3\n
// @lcpr case=end

*/

