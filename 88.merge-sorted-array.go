/*
 * @lc app=leetcode.cn id=88 lang=golang
 * @lcpr version=30204
 *
 * [88] 合并两个有序数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	t := m + n - 1
	for p2 >= 0 {
		if p1 < 0 || nums2[p2] >= nums1[p1] {
			nums1[t] = nums2[p2]
			t--
			p2--
		} else {
			nums1[t] = nums1[p1]
			t--
			p1--
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,0,0,0]\n3\n[2,5,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n[]\n0\n
// @lcpr case=end

// @lcpr case=start
// [0]\n0\n[1]\n1\n
// @lcpr case=end

*/

