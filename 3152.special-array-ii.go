/*
 * @lc app=leetcode.cn id=3152 lang=golang
 * @lcpr version=30204
 *
 * [3152] 特殊数组 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func isArraySpecial(nums []int, queries [][]int) []bool {
	pre := make([]int, len(nums)+1)
	for i, v := range nums {
		if i == 0 || (v+nums[i-1])%2 != 0 {
			pre[i+1] = 0
		}
	}

}

// @lc code=end

/*
// @lcpr case=start
// [3,4,1,2,6]\n[[0,4]]\n
// @lcpr case=end

// @lcpr case=start
// [4,3,1,6]\n[[0,2],[2,3]]\n
// @lcpr case=end

*/

