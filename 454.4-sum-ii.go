/*
 * @lc app=leetcode.cn id=454 lang=golang
 * @lcpr version=30204
 *
 * [454] 四数相加 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) (res int) {
	mp := map[int]int{}
	for _, i := range nums1 {
		for _, j := range nums2 {
			mp[i+j]++
		}
	}
	for _, i := range nums3 {
		for _, j := range nums4 {
			res += mp[-i-j]
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2]\n[-2,-1]\n[-1,2]\n[0,2]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n[0]\n[0]\n
// @lcpr case=end

*/

