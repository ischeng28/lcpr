/*
 * @lc app=leetcode.cn id=1818 lang=golang
 * @lcpr version=30204
 *
 * [1818] 绝对差值和
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	n := len(nums1)
	left, right := -1, n
	for left+1 < right {
		mid := left + (right-left)/2

	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,7,5]\n[2,3,5]\n
// @lcpr case=end

// @lcpr case=start
// [2,4,6,8,10]\n[2,4,6,8,10]\n
// @lcpr case=end

// @lcpr case=start
// [1,10,4,4,2,7]\n[9,3,5,1,7,4]\n
// @lcpr case=end

*/

