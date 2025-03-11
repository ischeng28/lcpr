/*
 * @lc app=leetcode.cn id=2570 lang=golang
 * @lcpr version=30204
 *
 * [2570] 合并两个二维数组 - 求和法
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	p1, m := 0, len(nums1)
	p2, n := 0, len(nums2)
	res := [][]int{}
	for p1 < m || p2 < n {
		if p1 == m {
			return append(res, nums2[p2:]...)
		}
		if p2 == n {
			return append(res, nums1[p1:]...)
		}
		if nums1[p1][0] == nums2[p2][0] {
			res = append(res, []int{nums1[p1][0], nums1[p1][1] + nums2[p2][1]})
			p1++
			p2++
		} else if nums1[p1][0] < nums2[p2][0] {
			res = append(res, nums1[p1])
			p1++
		} else {
			res = append(res, nums2[p2])
			p2++
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2],[2,3],[4,5]]\n[[1,4],[3,2],[4,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[2,4],[3,6],[5,5]]\n[[1,3],[4,3]]\n
// @lcpr case=end

*/

