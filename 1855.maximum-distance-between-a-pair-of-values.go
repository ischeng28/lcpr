/*
 * @lc app=leetcode.cn id=1855 lang=golang
 * @lcpr version=30204
 *
 * [1855] 下标对中的最大距离
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxDistance(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	res := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			res = max(res, j-i)
			j++
		} else {
			i++
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [55,30,5,4,2]\n[100,20,10,10,5]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2]\n[10,10,1]\n
// @lcpr case=end

// @lcpr case=start
// [30,29,19,5]\n[25,25,25,25,25]\n
// @lcpr case=end

*/

