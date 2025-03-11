/*
 * @lc app=leetcode.cn id=2540 lang=golang
 * @lcpr version=30204
 *
 * [2540] 最小公共值
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func getCommon(nums1 []int, nums2 []int) int {
	p1, p2 := 0, 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			return nums1[p1]
		}
		if nums1[p1] < nums2[p2] {
			p1++
		} else {
			p2++
		}

	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n[2,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,6]\n[2,3,4,5]\n
// @lcpr case=end

*/

